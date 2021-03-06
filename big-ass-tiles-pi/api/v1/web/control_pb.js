// source: control.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

goog.exportSymbol('proto.BigAssTilesPi.v1.control.ControlRequest', null, global);
goog.exportSymbol('proto.BigAssTilesPi.v1.control.EmptyResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.BigAssTilesPi.v1.control.ControlRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.BigAssTilesPi.v1.control.ControlRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.BigAssTilesPi.v1.control.ControlRequest.displayName = 'proto.BigAssTilesPi.v1.control.ControlRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.BigAssTilesPi.v1.control.EmptyResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.BigAssTilesPi.v1.control.EmptyResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.BigAssTilesPi.v1.control.EmptyResponse.displayName = 'proto.BigAssTilesPi.v1.control.EmptyResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.BigAssTilesPi.v1.control.ControlRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.BigAssTilesPi.v1.control.ControlRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.BigAssTilesPi.v1.control.ControlRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.BigAssTilesPi.v1.control.ControlRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    input: jspb.Message.getFieldWithDefault(msg, 1, 0),
    value: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.BigAssTilesPi.v1.control.ControlRequest}
 */
proto.BigAssTilesPi.v1.control.ControlRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.BigAssTilesPi.v1.control.ControlRequest;
  return proto.BigAssTilesPi.v1.control.ControlRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.BigAssTilesPi.v1.control.ControlRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.BigAssTilesPi.v1.control.ControlRequest}
 */
proto.BigAssTilesPi.v1.control.ControlRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setInput(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setValue(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.BigAssTilesPi.v1.control.ControlRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.BigAssTilesPi.v1.control.ControlRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.BigAssTilesPi.v1.control.ControlRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.BigAssTilesPi.v1.control.ControlRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInput();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getValue();
  if (f !== 0.0) {
    writer.writeDouble(
      2,
      f
    );
  }
};


/**
 * optional int32 input = 1;
 * @return {number}
 */
proto.BigAssTilesPi.v1.control.ControlRequest.prototype.getInput = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.BigAssTilesPi.v1.control.ControlRequest} returns this
 */
proto.BigAssTilesPi.v1.control.ControlRequest.prototype.setInput = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional double value = 2;
 * @return {number}
 */
proto.BigAssTilesPi.v1.control.ControlRequest.prototype.getValue = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.BigAssTilesPi.v1.control.ControlRequest} returns this
 */
proto.BigAssTilesPi.v1.control.ControlRequest.prototype.setValue = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.BigAssTilesPi.v1.control.EmptyResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.BigAssTilesPi.v1.control.EmptyResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.BigAssTilesPi.v1.control.EmptyResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.BigAssTilesPi.v1.control.EmptyResponse.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.BigAssTilesPi.v1.control.EmptyResponse}
 */
proto.BigAssTilesPi.v1.control.EmptyResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.BigAssTilesPi.v1.control.EmptyResponse;
  return proto.BigAssTilesPi.v1.control.EmptyResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.BigAssTilesPi.v1.control.EmptyResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.BigAssTilesPi.v1.control.EmptyResponse}
 */
proto.BigAssTilesPi.v1.control.EmptyResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.BigAssTilesPi.v1.control.EmptyResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.BigAssTilesPi.v1.control.EmptyResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.BigAssTilesPi.v1.control.EmptyResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.BigAssTilesPi.v1.control.EmptyResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


goog.object.extend(exports, proto.BigAssTilesPi.v1.control);
