package goiffel

import (
	"time"
        "encoding/json"
	"github.com/google/uuid"
)

func getTimeStampMilliSeconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func newEiffelEvent(typeName string, data interface{}, links []EiffelLink) EiffelEvent {
	return EiffelEvent{
		Meta: EiffelMeta {
			Id: uuid.New().String(),
			Type: typeName,
			Version: EventVersion,
			Time: getTimeStampMilliSeconds(),
		},
		Data: data,
		Links: links,
	}
}

func postParseEventData(parsed_data interface{}, evt *EiffelEvent) error {
	// We convert the map[string]interface{}
	// to json encoded data - then we unmarshal it
	// again
	temporary_json, err := json.Marshal(evt.Data)
	if err != nil { return err }


	err = json.Unmarshal(temporary_json, parsed_data)
	if err != nil { return err }

	evt.Data = parsed_data

	return nil
}

func postReceiveParser(evt *EiffelEvent) error {
	switch t := evt.Meta.Type; t {

	case EiffelArtifactCreatedEvent:
		var artifactCreated EiffelArtifactCreatedEventData
		return postParseEventData(&artifactCreated, evt)

	case EiffelArtifactPublishedEvent:
		var artifactPublished EiffelArtifactPublishedEventData
		return postParseEventData(&artifactPublished, evt)

	case EiffelCompositionDefinedEvent:
		var compositionDefined EiffelCompositionDefinedEventData
		return postParseEventData(&compositionDefined, evt)

	default:
		return nil
	}
}
