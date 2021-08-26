// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_SBQueue_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueue_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueueStatusARM, SBQueueStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueueStatusARM runs a test to see if a specific instance of SBQueue_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueueStatusARM(subject SBQueue_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueue_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SBQueue_StatusARM instances for property testing - lazily instantiated by SBQueueStatusARMGenerator()
var sbQueueStatusARMGenerator gopter.Gen

// SBQueueStatusARMGenerator returns a generator of SBQueue_StatusARM instances for property testing.
// We first initialize sbQueueStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBQueueStatusARMGenerator() gopter.Gen {
	if sbQueueStatusARMGenerator != nil {
		return sbQueueStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueStatusARM(generators)
	sbQueueStatusARMGenerator = gen.Struct(reflect.TypeOf(SBQueue_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueStatusARM(generators)
	AddRelatedPropertyGeneratorsForSBQueueStatusARM(generators)
	sbQueueStatusARMGenerator = gen.Struct(reflect.TypeOf(SBQueue_StatusARM{}), generators)

	return sbQueueStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSBQueueStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueueStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBQueueStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBQueueStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBQueuePropertiesStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_SBQueueProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueueProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueuePropertiesStatusARM, SBQueuePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueuePropertiesStatusARM runs a test to see if a specific instance of SBQueueProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueuePropertiesStatusARM(subject SBQueueProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueueProperties_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SBQueueProperties_StatusARM instances for property testing - lazily instantiated by
//SBQueuePropertiesStatusARMGenerator()
var sbQueuePropertiesStatusARMGenerator gopter.Gen

// SBQueuePropertiesStatusARMGenerator returns a generator of SBQueueProperties_StatusARM instances for property testing.
// We first initialize sbQueuePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBQueuePropertiesStatusARMGenerator() gopter.Gen {
	if sbQueuePropertiesStatusARMGenerator != nil {
		return sbQueuePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueuePropertiesStatusARM(generators)
	sbQueuePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SBQueueProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueuePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForSBQueuePropertiesStatusARM(generators)
	sbQueuePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SBQueueProperties_StatusARM{}), generators)

	return sbQueuePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSBQueuePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueuePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["MessageCount"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(EntityStatus_StatusActive, EntityStatus_StatusCreating, EntityStatus_StatusDeleting, EntityStatus_StatusDisabled, EntityStatus_StatusReceiveDisabled, EntityStatus_StatusRenaming, EntityStatus_StatusRestoring, EntityStatus_StatusSendDisabled, EntityStatus_StatusUnknown))
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBQueuePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBQueuePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetailsStatusARMGenerator())
}

func Test_MessageCountDetails_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MessageCountDetails_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMessageCountDetailsStatusARM, MessageCountDetailsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMessageCountDetailsStatusARM runs a test to see if a specific instance of MessageCountDetails_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMessageCountDetailsStatusARM(subject MessageCountDetails_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MessageCountDetails_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MessageCountDetails_StatusARM instances for property testing - lazily instantiated by
//MessageCountDetailsStatusARMGenerator()
var messageCountDetailsStatusARMGenerator gopter.Gen

// MessageCountDetailsStatusARMGenerator returns a generator of MessageCountDetails_StatusARM instances for property testing.
func MessageCountDetailsStatusARMGenerator() gopter.Gen {
	if messageCountDetailsStatusARMGenerator != nil {
		return messageCountDetailsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMessageCountDetailsStatusARM(generators)
	messageCountDetailsStatusARMGenerator = gen.Struct(reflect.TypeOf(MessageCountDetails_StatusARM{}), generators)

	return messageCountDetailsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForMessageCountDetailsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMessageCountDetailsStatusARM(gens map[string]gopter.Gen) {
	gens["ActiveMessageCount"] = gen.PtrOf(gen.Int())
	gens["DeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["ScheduledMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferDeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferMessageCount"] = gen.PtrOf(gen.Int())
}
