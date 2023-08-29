// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/check_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_relation_tuples_v1alpha2_check_service_pb from "../../../../ory/keto/relation_tuples/v1alpha2/check_service_pb";
import * as ory_keto_relation_tuples_v1alpha2_relation_tuples_pb from "../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb";

interface ICheckServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    check: ICheckServiceService_ICheck;
}

interface ICheckServiceService_ICheck extends grpc.MethodDefinition<ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse> {
    path: "/ory.keto.relation_tuples.v1alpha2.CheckService/Check";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest>;
    responseSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse>;
}

export const CheckServiceService: ICheckServiceService;

export interface ICheckServiceServer {
    check: grpc.handleUnaryCall<ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse>;
}

export interface ICheckServiceClient {
    check(request: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    check(request: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    check(request: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
}

export class CheckServiceClient extends grpc.Client implements ICheckServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public check(request: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    public check(request: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    public check(request: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
}
