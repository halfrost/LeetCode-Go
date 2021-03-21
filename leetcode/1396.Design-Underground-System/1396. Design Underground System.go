package leetcode

type checkin struct {
	station string
	time    int
}

type stationTime struct {
	sum, count float64
}

type UndergroundSystem struct {
	checkins     map[int]*checkin
	stationTimes map[string]map[string]*stationTime
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		make(map[int]*checkin),
		make(map[string]map[string]*stationTime),
	}
}

func (s *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	s.checkins[id] = &checkin{stationName, t}
}

func (s *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkin := s.checkins[id]
	destination := s.stationTimes[checkin.station]
	if destination == nil {
		s.stationTimes[checkin.station] = make(map[string]*stationTime)
	}
	st := s.stationTimes[checkin.station][stationName]
	if st == nil {
		st = new(stationTime)
		s.stationTimes[checkin.station][stationName] = st
	}
	st.sum += float64(t - checkin.time)
	st.count++
	delete(s.checkins, id)
}

func (s *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	st := s.stationTimes[startStation][endStation]
	return st.sum / st.count
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
