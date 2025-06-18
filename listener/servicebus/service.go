package servicebus

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-stater-listener/domain/model"
	"go-stater-listener/pkg/env"
	"go-stater-listener/pkg/utils"

	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/service_bus/topic"
)

type serviceBus struct {
	topic *topic.Topic
}

type ServiceBus interface {
	SubscriptionSuccess()
}

func NewServiceBus() ServiceBus {
	t := topic.New()
	// logM.Info(bpLogCenterModel.LoggerRequest{
	// 	Tags:    utils.TAGS_SERVICE_NEW_SERVICE,
	// 	Process: utils.PROCESS_START_PROJECT,
	// })
	return serviceBus{t}
}

func (s serviceBus) SubscriptionSuccess() {
	s.topic.SetTopic(env.Env().SB_TOPIC).SetSubscription(env.Env().SB_SUBSCRIPTION).Subscribe(s.Callback)
}

func (s serviceBus) CallProcess(message topic.MessageResponse) (als model.AppinsightLogStruct, errProp model.ErrorProps) {
	eventType, err := utils.EventType(message)
	if err != nil {
		errProp = utils.ErrorData(err)
		return
	}

	fmt.Println("")
	fmt.Println("eventType:", eventType)
	fmt.Println("")

	switch eventType {
	case "event_created":
		req := model.RequestEventCreate{}
		err = json.Unmarshal(message.Data, &req)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		// Simulate sending the event data to an external API
		url := fmt.Sprintf("%s/%s", env.Env().API_1, "events")
		eventData := map[string]interface{}{
			"eventName":      req.EventName,
			"price":          req.Price,
			"maxParticipant": req.MaxParticipant,
		}

		body, err := json.Marshal(eventData)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		header := http.Header{
			"Accept":       []string{"application/json;odata=verbose"},
			"Content-Type": []string{"application/json"},
		}

		client := &http.Client{}
		reqH, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			fmt.Println("Error creating request:", err)
		}

		reqH.Header = header

		res, err := client.Do(reqH)
		if err != nil {
			fmt.Println("Error making request:", err)
		}
		defer res.Body.Close()

		fmt.Println("Event created successfully")

	case "event_updated":
		req := model.RequestEventUpdated{}
		err = json.Unmarshal(message.Data, &req)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		// Simulate sending the event data to an external API
		url := fmt.Sprintf("%s/events/%d", env.Env().API_1, req.ID)
		eventData := map[string]interface{}{
			"eventName": req.EventName,
			"price":     req.Price,
		}

		body, err := json.Marshal(eventData)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		header := http.Header{
			"Accept":       []string{"application/json;odata=verbose"},
			"Content-Type": []string{"application/json"},
		}

		client := &http.Client{}
		reqH, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(body))
		if err != nil {
			fmt.Println("Error creating request:", err)
		}

		reqH.Header = header

		res, err := client.Do(reqH)
		if err != nil {
			fmt.Println("Error making request:", err)
		}
		defer res.Body.Close()

		fmt.Println("Event updated successfully")

	case "participant_enrolled":
		req := model.RequestParticipantEnrolled{}
		err = json.Unmarshal(message.Data, &req)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		// Simulate sending the event data to an external API
		url := fmt.Sprintf("%s/events/%d/enroll", env.Env().API_1, req.ID)
		eventData := map[string]interface{}{
			"participantIds": req.ParticipantIDs,
		}

		body, err := json.Marshal(eventData)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		header := http.Header{
			"Accept":       []string{"application/json;odata=verbose"},
			"Content-Type": []string{"application/json"},
		}

		client := &http.Client{}
		reqH, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			fmt.Println("Error creating request:", err)
		}

		reqH.Header = header

		res, err := client.Do(reqH)
		if err != nil {
			fmt.Println("Error making request:", err)
		}
		defer res.Body.Close()

		fmt.Println("Participant enrolled successfully")

	default:
		errProp = utils.ErrorData(errors.New("call process event type not found"))
		return
	}

	return als, errProp
}

func (s serviceBus) Callback(message topic.MessageResponse, receiver *azservicebus.Receiver, msg *azservicebus.ReceivedMessage) error {
	s.CallProcess(message)
	// defer s.HandlePanic(message)
	// als, errProp := s.CallProcess(message)
	// if errProp.Error != nil {
	// 	deadLetterOptions := &azservicebus.DeadLetterOptions{
	// 		ErrorDescription: to.Ptr(errProp.Error.Error()),
	// 	}
	// 	receiver.DeadLetterMessage(context.TODO(), msg, deadLetterOptions)

	// 	err := s.AppinsightMonitorLog(als, message, errProp)
	// 	if err != nil {
	// 		s.logM.Error(bpLogCenterModel.LoggerRequest{
	// 			Tags:    utils.TAGS_SERVICE_CALLBACK,
	// 			Process: utils.PROCESS_PANIC_LOG,
	// 			Error:   fmt.Sprintf("Error AppinsightMonitorLog: %v", err),
	// 		})
	// 	}
	// } else {
	// 	err := s.AppinsightMonitorLog(als, message, model.ErrorProps{})
	// 	if err != nil {
	// 		s.logM.Error(bpLogCenterModel.LoggerRequest{
	// 			Tags:    utils.TAGS_SERVICE_CALLBACK,
	// 			Process: utils.PROCESS_PANIC_LOG,
	// 			Error:   fmt.Sprintf("Error AppinsightMonitorLog: %v", err),
	// 		})
	// 	}
	// }

	err := receiver.CompleteMessage(context.Background(), msg, &azservicebus.CompleteMessageOptions{})
	if err != nil {
		// s.logM.Error(bpLogCenterModel.LoggerRequest{
		// 	Tags:    utils.TAGS_SERVICE_CALLBACK,
		// 	Process: utils.PROCESS_COMPLETE_ERROR,
		// 	Error:   fmt.Sprintf("[Subscription] Failed to complete message: %v", err),
		// })
		fmt.Println("[Subscription] Failed to complete message:", err)
	}
	return nil
}

// func (s serviceBus) HandlePanic(message topic.MessageResponse) {
// 	var msg map[string]interface{}
// 	err := json.Unmarshal(message.Data, &msg)
// 	if err != nil {
// 		s.logM.Error(bpLogCenterModel.LoggerRequest{
// 			Tags:    utils.TAGS_PANIC,
// 			Process: utils.PROCESS_PANIC_LOG,
// 			Error:   fmt.Sprintf("Error HandlePanic: %v", err),
// 		})
// 	}

// 	if r := recover(); r != nil {
// 		msgData := struct {
// 			Event   string
// 			Message interface{} `json:"payload"`
// 		}{
// 			Event:   "Panic",
// 			Message: msg,
// 		}

// 		panicData := struct {
// 			Recover    interface{}
// 			DebugStack string
// 		}{
// 			Recover:    r,
// 			DebugStack: string(debug.Stack()),
// 		}

// 		s.logM.Error(bpLogCenterModel.LoggerRequest{
// 			Tags:    utils.TAGS_PANIC,
// 			Process: utils.PROCESS_PANIC_LOG,
// 			Param:   msgData,
// 			Panic:   panicData,
// 		})
// 	}
// }

// func (s serviceBus) AppinsightMonitorLog(als model.AppinsightLogStruct, message topic.MessageResponse, errProp model.ErrorProps) error {
// 	var msg map[string]interface{}
// 	err := json.Unmarshal(message.Data, &msg)
// 	if err != nil {
// 		return err
// 	}

// 	if errProp.Error != nil {
// 		msgData := model.LogFormatStruct{
// 			Event:     "Exception",
// 			EventType: als.EventType,
// 			Source:    als.Source,
// 			Message:   msg,
// 		}

// 		errData := model.LogErrorStruct{
// 			MessageError: errProp.Error.Error(),
// 			FileError:    errProp.FileError,
// 			LineError:    errProp.LineError,
// 		}

// 		s.logM.Error(bpLogCenterModel.LoggerRequest{
// 			TxID:    als.TxID,
// 			Tags:    utils.TAGS_LOG,
// 			Process: utils.PROCESS_APP_LOG,
// 			Error:   errData,
// 			Param:   msgData,
// 		})

// 	} else {
// 		msgData := model.LogFormatStruct{
// 			Event:     als.Event,
// 			EventType: als.EventType,
// 			Source:    als.Source,
// 			Message:   msg,
// 		}

// 		switch {
// 		case utils.ContainsEventType(als.EventType, utils.EventTypeRequestStruct), utils.ContainsEventType(als.EventType, utils.EventTypeResponseStruct):
// 			s.logM.Info(bpLogCenterModel.LoggerRequest{
// 				TxID:    als.TxID,
// 				Tags:    utils.TAGS_LOG,
// 				Process: utils.PROCESS_APP_LOG,
// 				Param:   msgData,
// 			})

// 		default:
// 			return errors.New("appinsight monitor log event type not found")
// 		}
// 	}
// 	return nil
// }
