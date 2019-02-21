package entities.service;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.18.0)",
    comments = "Source: services/user.proto")
public final class UserGrpc {

  private UserGrpc() {}

  public static final String SERVICE_NAME = "service.User";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<entities.command.User.LogIn,
      entities.event.User.LoggedIn> getLogInMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LogIn",
      requestType = entities.command.User.LogIn.class,
      responseType = entities.event.User.LoggedIn.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<entities.command.User.LogIn,
      entities.event.User.LoggedIn> getLogInMethod() {
    io.grpc.MethodDescriptor<entities.command.User.LogIn, entities.event.User.LoggedIn> getLogInMethod;
    if ((getLogInMethod = UserGrpc.getLogInMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getLogInMethod = UserGrpc.getLogInMethod) == null) {
          UserGrpc.getLogInMethod = getLogInMethod = 
              io.grpc.MethodDescriptor.<entities.command.User.LogIn, entities.event.User.LoggedIn>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "service.User", "LogIn"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.command.User.LogIn.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.event.User.LoggedIn.getDefaultInstance()))
                  .setSchemaDescriptor(new UserMethodDescriptorSupplier("LogIn"))
                  .build();
          }
        }
     }
     return getLogInMethod;
  }

  private static volatile io.grpc.MethodDescriptor<entities.command.User.LogOut,
      entities.event.User.LoggedOut> getLogOutMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LogOut",
      requestType = entities.command.User.LogOut.class,
      responseType = entities.event.User.LoggedOut.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<entities.command.User.LogOut,
      entities.event.User.LoggedOut> getLogOutMethod() {
    io.grpc.MethodDescriptor<entities.command.User.LogOut, entities.event.User.LoggedOut> getLogOutMethod;
    if ((getLogOutMethod = UserGrpc.getLogOutMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getLogOutMethod = UserGrpc.getLogOutMethod) == null) {
          UserGrpc.getLogOutMethod = getLogOutMethod = 
              io.grpc.MethodDescriptor.<entities.command.User.LogOut, entities.event.User.LoggedOut>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "service.User", "LogOut"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.command.User.LogOut.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.event.User.LoggedOut.getDefaultInstance()))
                  .setSchemaDescriptor(new UserMethodDescriptorSupplier("LogOut"))
                  .build();
          }
        }
     }
     return getLogOutMethod;
  }

  private static volatile io.grpc.MethodDescriptor<entities.command.User.SignUp,
      entities.event.User.SignedUp> getSignUpMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SignUp",
      requestType = entities.command.User.SignUp.class,
      responseType = entities.event.User.SignedUp.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<entities.command.User.SignUp,
      entities.event.User.SignedUp> getSignUpMethod() {
    io.grpc.MethodDescriptor<entities.command.User.SignUp, entities.event.User.SignedUp> getSignUpMethod;
    if ((getSignUpMethod = UserGrpc.getSignUpMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getSignUpMethod = UserGrpc.getSignUpMethod) == null) {
          UserGrpc.getSignUpMethod = getSignUpMethod = 
              io.grpc.MethodDescriptor.<entities.command.User.SignUp, entities.event.User.SignedUp>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "service.User", "SignUp"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.command.User.SignUp.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.event.User.SignedUp.getDefaultInstance()))
                  .setSchemaDescriptor(new UserMethodDescriptorSupplier("SignUp"))
                  .build();
          }
        }
     }
     return getSignUpMethod;
  }

  private static volatile io.grpc.MethodDescriptor<entities.command.User.ChangePassword,
      entities.event.User.PasswordChanged> getChangePasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ChangePassword",
      requestType = entities.command.User.ChangePassword.class,
      responseType = entities.event.User.PasswordChanged.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<entities.command.User.ChangePassword,
      entities.event.User.PasswordChanged> getChangePasswordMethod() {
    io.grpc.MethodDescriptor<entities.command.User.ChangePassword, entities.event.User.PasswordChanged> getChangePasswordMethod;
    if ((getChangePasswordMethod = UserGrpc.getChangePasswordMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getChangePasswordMethod = UserGrpc.getChangePasswordMethod) == null) {
          UserGrpc.getChangePasswordMethod = getChangePasswordMethod = 
              io.grpc.MethodDescriptor.<entities.command.User.ChangePassword, entities.event.User.PasswordChanged>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "service.User", "ChangePassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.command.User.ChangePassword.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.event.User.PasswordChanged.getDefaultInstance()))
                  .setSchemaDescriptor(new UserMethodDescriptorSupplier("ChangePassword"))
                  .build();
          }
        }
     }
     return getChangePasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<entities.command.User.UpdateProfile,
      entities.event.User.ProfileUpdated> getUpdateProfileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateProfile",
      requestType = entities.command.User.UpdateProfile.class,
      responseType = entities.event.User.ProfileUpdated.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<entities.command.User.UpdateProfile,
      entities.event.User.ProfileUpdated> getUpdateProfileMethod() {
    io.grpc.MethodDescriptor<entities.command.User.UpdateProfile, entities.event.User.ProfileUpdated> getUpdateProfileMethod;
    if ((getUpdateProfileMethod = UserGrpc.getUpdateProfileMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getUpdateProfileMethod = UserGrpc.getUpdateProfileMethod) == null) {
          UserGrpc.getUpdateProfileMethod = getUpdateProfileMethod = 
              io.grpc.MethodDescriptor.<entities.command.User.UpdateProfile, entities.event.User.ProfileUpdated>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "service.User", "UpdateProfile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.command.User.UpdateProfile.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  entities.event.User.ProfileUpdated.getDefaultInstance()))
                  .setSchemaDescriptor(new UserMethodDescriptorSupplier("UpdateProfile"))
                  .build();
          }
        }
     }
     return getUpdateProfileMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserStub newStub(io.grpc.Channel channel) {
    return new UserStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new UserBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new UserFutureStub(channel);
  }

  /**
   */
  public static abstract class UserImplBase implements io.grpc.BindableService {

    /**
     */
    public void logIn(entities.command.User.LogIn request,
        io.grpc.stub.StreamObserver<entities.event.User.LoggedIn> responseObserver) {
      asyncUnimplementedUnaryCall(getLogInMethod(), responseObserver);
    }

    /**
     */
    public void logOut(entities.command.User.LogOut request,
        io.grpc.stub.StreamObserver<entities.event.User.LoggedOut> responseObserver) {
      asyncUnimplementedUnaryCall(getLogOutMethod(), responseObserver);
    }

    /**
     */
    public void signUp(entities.command.User.SignUp request,
        io.grpc.stub.StreamObserver<entities.event.User.SignedUp> responseObserver) {
      asyncUnimplementedUnaryCall(getSignUpMethod(), responseObserver);
    }

    /**
     */
    public void changePassword(entities.command.User.ChangePassword request,
        io.grpc.stub.StreamObserver<entities.event.User.PasswordChanged> responseObserver) {
      asyncUnimplementedUnaryCall(getChangePasswordMethod(), responseObserver);
    }

    /**
     */
    public void updateProfile(entities.command.User.UpdateProfile request,
        io.grpc.stub.StreamObserver<entities.event.User.ProfileUpdated> responseObserver) {
      asyncUnimplementedUnaryCall(getUpdateProfileMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getLogInMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                entities.command.User.LogIn,
                entities.event.User.LoggedIn>(
                  this, METHODID_LOG_IN)))
          .addMethod(
            getLogOutMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                entities.command.User.LogOut,
                entities.event.User.LoggedOut>(
                  this, METHODID_LOG_OUT)))
          .addMethod(
            getSignUpMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                entities.command.User.SignUp,
                entities.event.User.SignedUp>(
                  this, METHODID_SIGN_UP)))
          .addMethod(
            getChangePasswordMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                entities.command.User.ChangePassword,
                entities.event.User.PasswordChanged>(
                  this, METHODID_CHANGE_PASSWORD)))
          .addMethod(
            getUpdateProfileMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                entities.command.User.UpdateProfile,
                entities.event.User.ProfileUpdated>(
                  this, METHODID_UPDATE_PROFILE)))
          .build();
    }
  }

  /**
   */
  public static final class UserStub extends io.grpc.stub.AbstractStub<UserStub> {
    private UserStub(io.grpc.Channel channel) {
      super(channel);
    }

    private UserStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new UserStub(channel, callOptions);
    }

    /**
     */
    public void logIn(entities.command.User.LogIn request,
        io.grpc.stub.StreamObserver<entities.event.User.LoggedIn> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getLogInMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void logOut(entities.command.User.LogOut request,
        io.grpc.stub.StreamObserver<entities.event.User.LoggedOut> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getLogOutMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void signUp(entities.command.User.SignUp request,
        io.grpc.stub.StreamObserver<entities.event.User.SignedUp> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSignUpMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void changePassword(entities.command.User.ChangePassword request,
        io.grpc.stub.StreamObserver<entities.event.User.PasswordChanged> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getChangePasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateProfile(entities.command.User.UpdateProfile request,
        io.grpc.stub.StreamObserver<entities.event.User.ProfileUpdated> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getUpdateProfileMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class UserBlockingStub extends io.grpc.stub.AbstractStub<UserBlockingStub> {
    private UserBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private UserBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new UserBlockingStub(channel, callOptions);
    }

    /**
     */
    public entities.event.User.LoggedIn logIn(entities.command.User.LogIn request) {
      return blockingUnaryCall(
          getChannel(), getLogInMethod(), getCallOptions(), request);
    }

    /**
     */
    public entities.event.User.LoggedOut logOut(entities.command.User.LogOut request) {
      return blockingUnaryCall(
          getChannel(), getLogOutMethod(), getCallOptions(), request);
    }

    /**
     */
    public entities.event.User.SignedUp signUp(entities.command.User.SignUp request) {
      return blockingUnaryCall(
          getChannel(), getSignUpMethod(), getCallOptions(), request);
    }

    /**
     */
    public entities.event.User.PasswordChanged changePassword(entities.command.User.ChangePassword request) {
      return blockingUnaryCall(
          getChannel(), getChangePasswordMethod(), getCallOptions(), request);
    }

    /**
     */
    public entities.event.User.ProfileUpdated updateProfile(entities.command.User.UpdateProfile request) {
      return blockingUnaryCall(
          getChannel(), getUpdateProfileMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class UserFutureStub extends io.grpc.stub.AbstractStub<UserFutureStub> {
    private UserFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private UserFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new UserFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<entities.event.User.LoggedIn> logIn(
        entities.command.User.LogIn request) {
      return futureUnaryCall(
          getChannel().newCall(getLogInMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<entities.event.User.LoggedOut> logOut(
        entities.command.User.LogOut request) {
      return futureUnaryCall(
          getChannel().newCall(getLogOutMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<entities.event.User.SignedUp> signUp(
        entities.command.User.SignUp request) {
      return futureUnaryCall(
          getChannel().newCall(getSignUpMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<entities.event.User.PasswordChanged> changePassword(
        entities.command.User.ChangePassword request) {
      return futureUnaryCall(
          getChannel().newCall(getChangePasswordMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<entities.event.User.ProfileUpdated> updateProfile(
        entities.command.User.UpdateProfile request) {
      return futureUnaryCall(
          getChannel().newCall(getUpdateProfileMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LOG_IN = 0;
  private static final int METHODID_LOG_OUT = 1;
  private static final int METHODID_SIGN_UP = 2;
  private static final int METHODID_CHANGE_PASSWORD = 3;
  private static final int METHODID_UPDATE_PROFILE = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final UserImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(UserImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LOG_IN:
          serviceImpl.logIn((entities.command.User.LogIn) request,
              (io.grpc.stub.StreamObserver<entities.event.User.LoggedIn>) responseObserver);
          break;
        case METHODID_LOG_OUT:
          serviceImpl.logOut((entities.command.User.LogOut) request,
              (io.grpc.stub.StreamObserver<entities.event.User.LoggedOut>) responseObserver);
          break;
        case METHODID_SIGN_UP:
          serviceImpl.signUp((entities.command.User.SignUp) request,
              (io.grpc.stub.StreamObserver<entities.event.User.SignedUp>) responseObserver);
          break;
        case METHODID_CHANGE_PASSWORD:
          serviceImpl.changePassword((entities.command.User.ChangePassword) request,
              (io.grpc.stub.StreamObserver<entities.event.User.PasswordChanged>) responseObserver);
          break;
        case METHODID_UPDATE_PROFILE:
          serviceImpl.updateProfile((entities.command.User.UpdateProfile) request,
              (io.grpc.stub.StreamObserver<entities.event.User.ProfileUpdated>) responseObserver);
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

  private static abstract class UserBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return entities.service.UserOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("User");
    }
  }

  private static final class UserFileDescriptorSupplier
      extends UserBaseDescriptorSupplier {
    UserFileDescriptorSupplier() {}
  }

  private static final class UserMethodDescriptorSupplier
      extends UserBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    UserMethodDescriptorSupplier(String methodName) {
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
      synchronized (UserGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserFileDescriptorSupplier())
              .addMethod(getLogInMethod())
              .addMethod(getLogOutMethod())
              .addMethod(getSignUpMethod())
              .addMethod(getChangePasswordMethod())
              .addMethod(getUpdateProfileMethod())
              .build();
        }
      }
    }
    return result;
  }
}
