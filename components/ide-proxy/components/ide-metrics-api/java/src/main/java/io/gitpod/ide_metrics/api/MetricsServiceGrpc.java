// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package io.gitpod.ide_metrics.api;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.45.0)",
    comments = "Source: metrics.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class MetricsServiceGrpc {

  private MetricsServiceGrpc() {}

  public static final String SERVICE_NAME = "ide_metrics_api.MetricsService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest,
      io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> getReportMetricMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ReportMetric",
      requestType = io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest.class,
      responseType = io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest,
      io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> getReportMetricMethod() {
    io.grpc.MethodDescriptor<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest, io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> getReportMetricMethod;
    if ((getReportMetricMethod = MetricsServiceGrpc.getReportMetricMethod) == null) {
      synchronized (MetricsServiceGrpc.class) {
        if ((getReportMetricMethod = MetricsServiceGrpc.getReportMetricMethod) == null) {
          MetricsServiceGrpc.getReportMetricMethod = getReportMetricMethod =
              io.grpc.MethodDescriptor.<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest, io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ReportMetric"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MetricsServiceMethodDescriptorSupplier("ReportMetric"))
              .build();
        }
      }
    }
    return getReportMetricMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest,
      io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> getReportLogMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "reportLog",
      requestType = io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest.class,
      responseType = io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest,
      io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> getReportLogMethod() {
    io.grpc.MethodDescriptor<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest, io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> getReportLogMethod;
    if ((getReportLogMethod = MetricsServiceGrpc.getReportLogMethod) == null) {
      synchronized (MetricsServiceGrpc.class) {
        if ((getReportLogMethod = MetricsServiceGrpc.getReportLogMethod) == null) {
          MetricsServiceGrpc.getReportLogMethod = getReportLogMethod =
              io.grpc.MethodDescriptor.<io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest, io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "reportLog"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MetricsServiceMethodDescriptorSupplier("reportLog"))
              .build();
        }
      }
    }
    return getReportLogMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static MetricsServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MetricsServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MetricsServiceStub>() {
        @java.lang.Override
        public MetricsServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MetricsServiceStub(channel, callOptions);
        }
      };
    return MetricsServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static MetricsServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MetricsServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MetricsServiceBlockingStub>() {
        @java.lang.Override
        public MetricsServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MetricsServiceBlockingStub(channel, callOptions);
        }
      };
    return MetricsServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static MetricsServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MetricsServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MetricsServiceFutureStub>() {
        @java.lang.Override
        public MetricsServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MetricsServiceFutureStub(channel, callOptions);
        }
      };
    return MetricsServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class MetricsServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void reportMetric(io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request,
        io.grpc.stub.StreamObserver<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReportMetricMethod(), responseObserver);
    }

    /**
     */
    public void reportLog(io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request,
        io.grpc.stub.StreamObserver<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReportLogMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getReportMetricMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest,
                io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse>(
                  this, METHODID_REPORT_METRIC)))
          .addMethod(
            getReportLogMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest,
                io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse>(
                  this, METHODID_REPORT_LOG)))
          .build();
    }
  }

  /**
   */
  public static final class MetricsServiceStub extends io.grpc.stub.AbstractAsyncStub<MetricsServiceStub> {
    private MetricsServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MetricsServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MetricsServiceStub(channel, callOptions);
    }

    /**
     */
    public void reportMetric(io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request,
        io.grpc.stub.StreamObserver<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReportMetricMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void reportLog(io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request,
        io.grpc.stub.StreamObserver<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReportLogMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class MetricsServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<MetricsServiceBlockingStub> {
    private MetricsServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MetricsServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MetricsServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse reportMetric(io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReportMetricMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse reportLog(io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReportLogMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class MetricsServiceFutureStub extends io.grpc.stub.AbstractFutureStub<MetricsServiceFutureStub> {
    private MetricsServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MetricsServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MetricsServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> reportMetric(
        io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReportMetricMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse> reportLog(
        io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReportLogMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_REPORT_METRIC = 0;
  private static final int METHODID_REPORT_LOG = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final MetricsServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(MetricsServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_REPORT_METRIC:
          serviceImpl.reportMetric((io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest) request,
              (io.grpc.stub.StreamObserver<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse>) responseObserver);
          break;
        case METHODID_REPORT_LOG:
          serviceImpl.reportLog((io.gitpod.ide_metrics.api.Metrics.ReportMetricsRequest) request,
              (io.grpc.stub.StreamObserver<io.gitpod.ide_metrics.api.Metrics.ReportMetricsResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class MetricsServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    MetricsServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.gitpod.ide_metrics.api.Metrics.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("MetricsService");
    }
  }

  private static final class MetricsServiceFileDescriptorSupplier
      extends MetricsServiceBaseDescriptorSupplier {
    MetricsServiceFileDescriptorSupplier() {}
  }

  private static final class MetricsServiceMethodDescriptorSupplier
      extends MetricsServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    MetricsServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (MetricsServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new MetricsServiceFileDescriptorSupplier())
              .addMethod(getReportMetricMethod())
              .addMethod(getReportLogMethod())
              .build();
        }
      }
    }
    return result;
  }
}
