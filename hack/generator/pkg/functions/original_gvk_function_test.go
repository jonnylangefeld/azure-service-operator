/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func Test_OriginalGVKFunction_ReadingOriginalVersionFromProperty_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup := "microsoft.person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	originalGVKFunction := NewOriginalGVKFunction(ReadProperty, idFactory)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status, originalGVKFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "OriginalGVKFunction", resource)
}

func Test_OriginalGVKFunction_ReadingOriginalVersionFromFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup := "microsoft.person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	originalGVKFunction := NewOriginalGVKFunction(ReadFunction, idFactory)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status, originalGVKFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "OriginalGVKFunction", resource)
}
