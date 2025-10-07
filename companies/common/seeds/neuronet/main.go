package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const ReynaRoute = "C:\\Users\\altte\\OneDrive\\Desktop\\GoPackMVP2025\\companies\\common\\seeds\\reyna\\reyna_route.json"

// Структура для одного города
type City struct {
	CityNum       int     `json:"-"`
	CityName      string  `json:"name"`
	TimeArrival   string  `json:"timeArrive"`
	Stand         string  `json:"stand"`
	TimeDeparture string  `json:"timeDepart"`
	Latitude      float64 `json:"lat,omitempty"`
	Longitude     float64 `json:"lon,omitempty"`
}

// Структура для всего маршрута
type RouteData map[string]City

// Обёртка для удобной работы
type Route struct {
	Cities []City
}

// Расширенная информация о часовом поясе
type TimeZoneInfo struct {
	CityName           string
	CurrentLocalTime   string
	MoscowTime         string
	TimeDiff           int
	IsMoving           bool
	NextStation        string
	DayOfTravel        int
	DistanceFromMoscow float64
	EstimatedDelay     time.Duration
}

// Реальная база часовых поясов для городов России
var russianTimeZones = map[string]int{
	"Москва":          0,
	"Санкт-Петербург": 0,
	"Владимир":        0,
	"Владимир Пасс":   0,
	"Нижний Новгород": 0,
	"Казань":          0,
	"Пермь":           2,
	"Екатеринбург":    2,
	"Тюмень":          2,
	"Омск":            3,
	"Новосибирск":     4,
	"Красноярск":      4,
	"Иркутск":         5,
	"Улан-Удэ":        5,
	"Чита":            6,
	"Хабаровск":       7,
	"Владивосток":     7,
}

// База координат городов (упрощённо)
var cityCoordinates = map[string]struct {
	Lat float64
	Lon float64
}{
	"Москва":        {55.7558, 37.6173},
	"Владимир Пасс": {56.1290, 40.4066},
	"Хабаровск":     {48.4827, 135.0840},
}

// Парсинг JSON
func parseRouteJSON(filename string) (*Route, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var routeData RouteData
	if err := json.Unmarshal(data, &routeData); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	// Конвертируем map в slice и добавляем номера городов
	var cities []City
	for key, city := range routeData {
		if num, err := strconv.Atoi(strings.TrimPrefix(key, "city_")); err == nil {
			city.CityNum = num
		}

		// Добавляем координаты если есть
		if coords, exists := cityCoordinates[city.CityName]; exists {
			city.Latitude = coords.Lat
			city.Longitude = coords.Lon
		}

		cities = append(cities, city)
	}

	// Сортируем по номерам городов
	sort.Slice(cities, func(i, j int) bool {
		return cities[i].CityNum < cities[j].CityNum
	})

	return &Route{Cities: cities}, nil
}

// Анализатор маршрута с расширенной функциональностью
type RouteAnalyzer struct {
	route      *Route
	moscowTime time.Time
	currentDay int
	workerPool chan struct{}
	mu         sync.RWMutex
	delays     map[string]time.Duration
}

func NewRouteAnalyzer(route *Route, moscowTimeStr string) (*RouteAnalyzer, error) {
	parsedTime, err := time.Parse("15:04", moscowTimeStr)
	if err != nil {
		return nil, fmt.Errorf("неверный формат времени: %v", err)
	}

	now := time.Now()
	moscowTime := time.Date(now.Year(), now.Month(), now.Day(),
		parsedTime.Hour(), parsedTime.Minute(), 0, 0, time.UTC)

	return &RouteAnalyzer{
		route:      route,
		moscowTime: moscowTime,
		workerPool: make(chan struct{}, 5), // Пул из 5 воркеров
		delays:     make(map[string]time.Duration),
		currentDay: 7, // 7 октября - начало маршрута
	}, nil
}

// Воркер для параллельного анализа сегментов маршрута
func (ra *RouteAnalyzer) analyzeSegmentWorker(segmentStart int, results chan<- *TimeZoneInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	ra.workerPool <- struct{}{}        // Занимаем воркера
	defer func() { <-ra.workerPool }() // Освобождаем воркера

	for i := segmentStart; i < len(ra.route.Cities)-1 && i < segmentStart+3; i++ {
		if info := ra.checkCityPosition(i); info != nil {
			results <- info
			return
		}
	}
}

// Алгоритм двух указателей для поиска текущего положения
func (ra *RouteAnalyzer) findWithTwoPointers() *TimeZoneInfo {
	left, right := 0, len(ra.route.Cities)-1

	for left <= right {
		mid := (left + right) / 2
		info := ra.checkCityPosition(mid)

		if info != nil {
			return info
		}

		// Определяем направление поиска
		midArrival, _ := time.Parse("15:04", ra.route.Cities[mid].TimeArrival)
		midArrivalNorm := time.Date(ra.moscowTime.Year(), ra.moscowTime.Month(), ra.moscowTime.Day(),
			midArrival.Hour(), midArrival.Minute(), 0, 0, time.UTC)

		if ra.moscowTime.Before(midArrivalNorm) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nil
}

// Алгоритм скользящего окна для прогнозирования
func (ra *RouteAnalyzer) slidingWindowPrediction() *TimeZoneInfo {
	windowSize := 3
	if len(ra.route.Cities) < windowSize {
		windowSize = len(ra.route.Cities)
	}

	for i := 0; i <= len(ra.route.Cities)-windowSize; i++ {
		window := ra.route.Cities[i : i+windowSize]

		for j := 0; j < len(window)-1; j++ {
			if info := ra.checkCityPosition(i + j); info != nil {
				// Добавляем прогноз задержки
				info.EstimatedDelay = ra.calculateDelay(i + j)
				return info
			}
		}
	}
	return nil
}

// Хэш-мап поиск (быстрый поиск по городу)
func (ra *RouteAnalyzer) hashMapSearch() *TimeZoneInfo {
	cityTimeMap := make(map[string]time.Time)

	for _, city := range ra.route.Cities {
		arrival, _ := time.Parse("15:04", city.TimeArrival)
		cityTimeMap[city.CityName] = time.Date(ra.moscowTime.Year(), ra.moscowTime.Month(), ra.moscowTime.Day(),
			arrival.Hour(), arrival.Minute(), 0, 0, time.UTC)
	}

	// Ищем ближайший город по времени
	var closestCity City
	var minDiff time.Duration

	for _, city := range ra.route.Cities {
		cityTime := cityTimeMap[city.CityName]
		diff := ra.moscowTime.Sub(cityTime).Abs()

		if minDiff == 0 || diff < minDiff {
			minDiff = diff
			closestCity = city
		}
	}

	return ra.createDetailedTimeZoneInfo(closestCity, "", false)
}

// Основной метод определения положения с выбором алгоритма
func (ra *RouteAnalyzer) FindCurrentPosition() *TimeZoneInfo {
	// Пробуем разные алгоритмы
	if result := ra.findWithTwoPointers(); result != nil {
		return result
	}

	if result := ra.slidingWindowPrediction(); result != nil {
		return result
	}

	return ra.hashMapSearch()
}

// Расчёт задержки относительно графика
func (ra *RouteAnalyzer) calculateDelay(cityIndex int) time.Duration {
	if cityIndex == 0 {
		return 0
	}

	// Средняя скорость поезда 80 км/ч
	averageSpeed := 80.0
	expectedTravelTime := ra.calculateExpectedTravelTime(cityIndex-1, cityIndex, averageSpeed)

	// Здесь можно добавить логику расчёта реальной задержки
	// Пока возвращаем нулевую задержку
	return 0
}

// Расчёт ожидаемого времени пути между городами
func (ra *RouteAnalyzer) calculateExpectedTravelTime(from, to int, speed float64) time.Duration {
	if from >= len(ra.route.Cities) || to >= len(ra.route.Cities) {
		return 0
	}

	// Упрощённый расчёт расстояния по координатам
	distance := ra.calculateDistance(ra.route.Cities[from], ra.route.Cities[to])
	hours := distance / speed
	return time.Duration(hours * float64(time.Hour))
}

// Расчёт расстояния между городами (упрощённая формула)
func (ra *RouteAnalyzer) calculateDistance(city1, city2 City) float64 {
	if city1.Latitude == 0 || city2.Latitude == 0 {
		// Если координат нет, используем приблизительные расстояния
		distances := map[string]float64{
			"Москва-Владимир Пасс":    180,
			"Владимир Пасс-Хабаровск": 8300,
		}
		key := city1.CityName + "-" + city2.CityName
		if dist, exists := distances[key]; exists {
			return dist
		}
		return 100 // км по умолчанию
	}

	// Упрощённый расчёт расстояния
	latDiff := city2.Latitude - city1.Latitude
	lonDiff := city2.Longitude - city1.Longitude
	return (latDiff*latDiff + lonDiff*lonDiff) * 110 // приблизительно в км
}

// Определение дня путешеления
func (ra *RouteAnalyzer) calculateDayOfTravel(cityIndex int) int {
	baseDay := 7 // 7 октября
	if cityIndex == 0 {
		return baseDay
	}

	// Логика определения дня на основе времени прибытия
	// Упрощённо - каждый город через определённое время
	hoursTraveled := cityIndex * 3 // приблизительно
	daysPassed := hoursTraveled / 24

	return baseDay + daysPassed
}

// Расстояние от Москвы
func (ra *RouteAnalyzer) calculateDistanceFromMoscow(cityIndex int) float64 {
	totalDistance := 0.0
	for i := 0; i < cityIndex && i < len(ra.route.Cities)-1; i++ {
		totalDistance += ra.calculateDistance(ra.route.Cities[i], ra.route.Cities[i+1])
	}
	return totalDistance
}

func (ra *RouteAnalyzer) checkCityPosition(i int) *TimeZoneInfo {
	if i >= len(ra.route.Cities)-1 {
		return nil
	}

	currentCity := ra.route.Cities[i]
	nextCity := ra.route.Cities[i+1]

	arrivalTime, _ := time.Parse("15:04", currentCity.TimeArrival)
	departureTime, _ := time.Parse("15:04", currentCity.TimeDeparture)
	nextArrival, _ := time.Parse("15:04", nextCity.TimeArrival)

	currentArrival := time.Date(ra.moscowTime.Year(), ra.moscowTime.Month(), ra.moscowTime.Day(),
		arrivalTime.Hour(), arrivalTime.Minute(), 0, 0, time.UTC)
	currentDeparture := time.Date(ra.moscowTime.Year(), ra.moscowTime.Month(), ra.moscowTime.Day(),
		departureTime.Hour(), departureTime.Minute(), 0, 0, time.UTC)
	nextCityArrival := time.Date(ra.moscowTime.Year(), ra.moscowTime.Month(), ra.moscowTime.Day(),
		nextArrival.Hour(), nextArrival.Minute(), 0, 0, time.UTC)

	if ra.moscowTime.After(currentArrival) && ra.moscowTime.Before(currentDeparture) {
		return ra.createDetailedTimeZoneInfo(currentCity, nextCity.CityName, false)
	}

	if ra.moscowTime.After(currentDeparture) && ra.moscowTime.Before(nextCityArrival) {
		return ra.createDetailedTimeZoneInfo(currentCity, nextCity.CityName, true)
	}

	return nil
}

func (ra *RouteAnalyzer) getTimeZoneOffset(cityName string) int {
	if offset, exists := russianTimeZones[cityName]; exists {
		return offset
	}
	return 0
}

func (ra *RouteAnalyzer) createDetailedTimeZoneInfo(city City, nextStation string, isMoving bool) *TimeZoneInfo {
	timeDiff := ra.getTimeZoneOffset(city.CityName)
	localTime := ra.moscowTime.Add(time.Duration(timeDiff) * time.Hour)
	dayOfTravel := ra.calculateDayOfTravel(city.CityNum - 1)
	distanceFromMoscow := ra.calculateDistanceFromMoscow(city.CityNum - 1)

	return &TimeZoneInfo{
		CityName:           city.CityName,
		CurrentLocalTime:   localTime.Format("15:04"),
		MoscowTime:         ra.moscowTime.Format("15:04"),
		TimeDiff:           timeDiff,
		IsMoving:           isMoving,
		NextStation:        nextStation,
		DayOfTravel:        dayOfTravel,
		DistanceFromMoscow: distanceFromMoscow,
	}
}

// Генерация расписания по дням
func (ra *RouteAnalyzer) GenerateDailySchedule() map[int][]City {
	dailySchedule := make(map[int][]City)

	for _, city := range ra.route.Cities {
		day := ra.calculateDayOfTravel(city.CityNum - 1)
		dailySchedule[day] = append(dailySchedule[day], city)
	}

	return dailySchedule
}

// Финальный вывод с расширенной информацией
func verdict(analyzer *RouteAnalyzer) string {
	position := analyzer.FindCurrentPosition()
	if position == nil {
		return "Не удалось определить положение поезда"
	}

	var status string
	if position.IsMoving {
		status = fmt.Sprintf("🚂 Поезд в пути до станции: %s", position.NextStation)
	} else {
		status = fmt.Sprintf("🛑 Поезд на станции: %s", position.CityName)
	}

	return fmt.Sprintf(`
%s

📅 День путешествия: %d октября
⏰ Время в Москве: %s
🌍 Локальное время: %s
📊 Разница во времени: +%d часов
📏 Расстояние от Москвы: %.0f км
	`, status, position.DayOfTravel, position.MoscowTime,
		position.CurrentLocalTime, position.TimeDiff, position.DistanceFromMoscow)
}

func main() {
	// Парсим маршрут
	route, err := parseRouteJSON(ReynaRoute)
	if err != nil {
		log.Fatalf("Ошибка загрузки маршрута: %v", err)
	}

	// Получаем время от пользователя
	var moscowTime string
	fmt.Println("Какое время сейчас в Москве? Формат: 11:11")
	fmt.Scan(&moscowTime)

	// Создаём анализатор
	analyzer, err := NewRouteAnalyzer(route, moscowTime)
	if err != nil {
		log.Fatalf("Ошибка создания анализатора: %v", err)
	}

	// Получаем результат
	result := verdict(analyzer)
	fmt.Println(result)

	// Генерируем расписание по дням
	fmt.Println("\n📅 Расписание по дням:")
	dailySchedule := analyzer.GenerateDailySchedule()
	for day := 7; day <= 13; day++ {
		if cities, exists := dailySchedule[day]; exists {
			fmt.Printf("\n%d октября:\n", day)
			for _, city := range cities {
				fmt.Printf("  🏙️  %s - Прибытие: %s, Отправление: %s\n",
					city.CityName, city.TimeArrival, city.TimeDeparture)
			}
		}
	}

	// Весь маршрут
	fmt.Println("\n📋 Полный маршрут:")
	for _, city := range route.Cities {
		timeDiff := analyzer.getTimeZoneOffset(city.CityName)
		fmt.Printf("Город %d: %s (UTC+%d) - Прибытие: %s, Стоянка: %s, Отправление: %s\n",
			city.CityNum, city.CityName, timeDiff, city.TimeArrival, city.Stand, city.TimeDeparture)
	}
}
