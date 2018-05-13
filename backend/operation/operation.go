package riderOps

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"test-fullstack-loyalty/backend/model"
	"test-fullstack-loyalty/backend/store"
)

type Operation struct {
	store store.Store
}

func NewOperation(store store.Store) *Operation {
	if store == nil {
		log.Printf("Fail to create Operation due to empty store")
		return nil
	}
	return &Operation{store}
}

// GetUser get rider info with riderId
//func (ops *Operation) CUser(key string) (*model.Rider, error) {
//	var rider = new(model.Rider)
//	if err := redis.ScanStruct(value, rider); err != nil {
//		return nil, err
//	}
//	return rider, nil
//}

// genUserKey retrieves user based on riderId
//func (ops *Operation) getRiderKey(riderId string) (string, error) {
//var riderIds []string
//var err error
//formatRiderId := fmt.Sprintf("%s:*:%s", riderId, "user")
//if riderIds, err = ops.store.Keys(formatRiderId); err != nil {
//	log.Printf("Fail to get key %s, %v", riderIds, err)
//}
//if len(riderIds) != 1 {
//	return "", fmt.Errorf("More/less than one rider found  %v", riderIds)
//}
//return riderIds[0], nil
//}

// CreateInterviewer create rider with key riderId:timestamp:key
func (ops *Operation) CreateInterviewer(data model.Interviewer) (string, error) {
	id := fmt.Printf("%s", time.Now().Unix())
	newInterviewer := &model.Rider{
		Id:       id,
		Name:     data.Name,
		TimeSlot: make(model.TimeSlot, 0),
	}
	if err := ops.store.Set(id, newInterviewer); err != nil {
		return "", err
	}
	return id, nil
}

// createRide creates a ride
//func (ops *RiderOperation) CreateRide(data model.Payload) (string, error) {
//
//	defer monitoring.RideCreatedReceived.Inc()
//
//	var err error
//	// key is in the format of riderId:timestamp:created
//	key := genKey(data.RiderId, "created")
//	if err = ops.newRide(data, key); err != nil {
//		log.Printf("[Error] Fail to create a ride, %v", err)
//	} else {
//		monitoring.RideCreated.Inc()
//	}
//	return key, ops.newRide(data, key)
//}
//
//// CompleteRide completes a ride, update loyalty and push to frontend
//func (ops *RiderOperation) CompleteRide(data model.Payload) (string, error) {
//
//	defer monitoring.RideCompleteReceived.Inc()
//
//	var err error
//	var rider *model.Rider
//
//	// key is in the format of riderId:timestamp:complete
//	key := genKey(data.RiderId, "complete")
//	if err = ops.newRide(data, key); err != nil {
//		return "", err
//	}
//
//	// Update loyalty
//	var keyStr = strconv.FormatInt(data.RiderId, 10)
//	if rider, err = ops.updateLoyalty(keyStr, data.Amount, loyaltyCalculation); err != nil {
//		log.Printf("[Error] Fail to update loyalty %v", err)
//		return "", err
//	}
//
//	monitoring.RideComplete.Inc()
//
//	// Live update rider info to the frontend
//	ops.riderPusher <- *rider
//	return key, nil
//}
//
//// newRide writes a created/complete ride record in redis
//func (ops *RiderOperation) newRide(data model.Payload, key string) error {
//	ride := &model.Ride{
//		Id:      data.Id,
//		Amount:  data.Amount,
//		RiderId: data.RiderId,
//	}
//	if err := ops.store.Set(key, ride); err != nil {
//		log.Printf("[Error] Fail to create/complete new ride %v, %v", err, data)
//		return err
//	}
//	return nil
//}
//
//// loyaltyCalculation update loyalty logic
//func loyaltyCalculation(rider model.Rider, amount float32) model.Rider {
//	rider.NumRides = rider.NumRides + 1
//	if rider.NumRides < 20 {
//		rider.Loyalty = rider.Loyalty + amount
//	} else if rider.NumRides < 50 {
//		rider.Loyalty = rider.Loyalty + amount*3
//		if rider.NumRides == 20 {
//			rider.Grade = "SILVER"
//			monitoring.NumSilverRider.Inc()
//			monitoring.NumBronzeRider.Dec()
//		}
//	} else if rider.NumRides < 100 {
//		rider.Loyalty = rider.Loyalty + amount*5
//		if rider.NumRides == 50 {
//			rider.Grade = "GOLD"
//			monitoring.NumGoldRider.Inc()
//			monitoring.NumSilverRider.Dec()
//		}
//	} else {
//		rider.Loyalty = rider.Loyalty + amount*10
//		if rider.NumRides == 100 {
//			rider.Grade = "PLATINUM"
//			monitoring.NumPlatinumRider.Inc()
//			monitoring.NumGoldRider.Dec()
//		}
//	}
//	return rider
//}
//
//// updateLoyalty updates loyalty, grade and num of rides
//func (ops *RiderOperation) updateLoyalty(key string, amount float32, updateRider LoyaltyUpdateFunc) (*model.Rider, error) {
//	rider, err := ops.GetRider(key)
//	if err != nil {
//		log.Printf("Fail to get Rider with key %s", key)
//		return nil, err
//	}
//
//	updatedRider := updateRider(*rider, amount)
//
//	// Update loyalty
//	var riderIdStr = strconv.FormatInt(updatedRider.Id, 10)
//	riderKey, _ := ops.getRiderKey(riderIdStr)
//	if err := ops.store.Set(riderKey, updatedRider); err != nil {
//		return nil, err
//	}
//	return &updatedRider, nil
//}
//
// genKey generates key as riderId:timestamp:suffix
func genKey(key int64) string {
	var keyStr = strconv.FormatInt(key, 10)
	var timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	return fmt.Sprintf("%s:%s", keyStr, timestamp)
}
