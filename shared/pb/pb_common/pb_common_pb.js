// source: pb_common/pb_common.proto
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
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

goog.exportSymbol('proto.pb_common.ROUTE', null, global);
/**
 * @enum {number}
 */
proto.pb_common.ROUTE = {
  ROUTE_NONE: 0,
  LOBBY_ECHO: 1001,
  LOBBY_LOGIN: 1002,
  MATCH_ECHO: 2001,
  MATCH_JOIN: 2002,
  RPSGAME_ECHO: 3001,
  RPSGAME_START: 3002,
  RPSGAME_PLAY: 3003,
  RPSGAME_RESULT: 3004
};

goog.object.extend(exports, proto.pb_common);
