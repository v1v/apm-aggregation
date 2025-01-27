// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

// Code generated by protohash/generate. DO NOT EDIT.

package protohash

import (
	"encoding/binary"

	"github.com/cespare/xxhash/v2"

	"github.com/elastic/apm-aggregation/aggregationpb"
)

func writeUint32(h *xxhash.Digest, v uint32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], v)
	h.Write(buf[:])
}

func writeUint64(h *xxhash.Digest, v uint64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], v)
	h.Write(buf[:])
}

func HashServiceAggregationKey(h xxhash.Digest, k *aggregationpb.ServiceAggregationKey) xxhash.Digest {
	writeUint64(&h, k.Timestamp)
	h.WriteString(k.ServiceName)
	h.WriteString(k.ServiceEnvironment)
	h.WriteString(k.ServiceLanguageName)
	h.WriteString(k.AgentName)
	return h
}

func HashServiceInstanceAggregationKey(h xxhash.Digest, k *aggregationpb.ServiceInstanceAggregationKey) xxhash.Digest {
	h.Write(k.GlobalLabelsStr)
	return h
}

func HashServiceTransactionAggregationKey(h xxhash.Digest, k *aggregationpb.ServiceTransactionAggregationKey) xxhash.Digest {
	h.WriteString(k.TransactionType)
	return h
}

func HashSpanAggregationKey(h xxhash.Digest, k *aggregationpb.SpanAggregationKey) xxhash.Digest {
	h.WriteString(k.SpanName)
	h.WriteString(k.Outcome)
	h.WriteString(k.TargetType)
	h.WriteString(k.TargetName)
	h.WriteString(k.Resource)
	return h
}

func HashTransactionAggregationKey(h xxhash.Digest, k *aggregationpb.TransactionAggregationKey) xxhash.Digest {
	if k.TraceRoot {
		h.WriteString("1")
	}
	h.WriteString(k.ContainerId)
	h.WriteString(k.KubernetesPodName)
	h.WriteString(k.ServiceVersion)
	h.WriteString(k.ServiceNodeName)
	h.WriteString(k.ServiceRuntimeName)
	h.WriteString(k.ServiceRuntimeVersion)
	h.WriteString(k.ServiceLanguageVersion)
	h.WriteString(k.HostHostname)
	h.WriteString(k.HostName)
	h.WriteString(k.HostOsPlatform)
	h.WriteString(k.EventOutcome)
	h.WriteString(k.TransactionName)
	h.WriteString(k.TransactionType)
	h.WriteString(k.TransactionResult)
	writeUint32(&h, k.FaasColdstart)
	h.WriteString(k.FaasId)
	h.WriteString(k.FaasName)
	h.WriteString(k.FaasVersion)
	h.WriteString(k.FaasTriggerType)
	h.WriteString(k.CloudProvider)
	h.WriteString(k.CloudRegion)
	h.WriteString(k.CloudAvailabilityZone)
	h.WriteString(k.CloudServiceName)
	h.WriteString(k.CloudAccountId)
	h.WriteString(k.CloudAccountName)
	h.WriteString(k.CloudMachineType)
	h.WriteString(k.CloudProjectId)
	h.WriteString(k.CloudProjectName)
	return h
}
