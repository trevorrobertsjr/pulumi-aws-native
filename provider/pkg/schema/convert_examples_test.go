// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamoKeySchemaConversion(t *testing.T) {
	input := map[string]interface{}{
		"keySchema": []interface{}{
			map[string]interface{}{
				"attributeName": "test",
				"keyType":       "HASH",
			},
			map[string]interface{}{
				"attributeName": "Part",
				"keyType":       "RANGE",
			},
		},
		"sseSpecification": map[string]interface{}{
			"kmsMasterKeyId": "key",
			"sseEnabled":     true,
		},
	}

	expected := map[string]interface{}{
		"KeySchema": []interface{}{
			map[string]interface{}{
				"AttributeName": "test",
				"KeyType":       "HASH",
			},
			map[string]interface{}{
				"AttributeName": "Part",
				"KeyType":       "RANGE",
			},
		},
		"SSESpecification": map[string]interface{}{
			"KMSMasterKeyId": "key",
			"SSEEnabled":     true,
		},
	}
	testConversion(t, "aws-native:dynamodb:Table", input, expected)
}

func testConversion(t *testing.T, resource string, input map[string]interface{}, expected map[string]interface{}) {
	bytes, err := ioutil.ReadFile("../../cmd/pulumi-resource-aws-native/metadata.json")
	assert.NoError(t, err)
	metadata := CloudAPIMetadata{}
	err = json.Unmarshal(bytes, &metadata)
	assert.NoError(t, err)
	res := metadata.Resources[resource]
	actual, err := SdkToCfn(&res, metadata.Types, input)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
