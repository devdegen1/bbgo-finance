# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: bbgo.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nbbgo.proto\x12\x02pb\"\x07\n\x05\x45mpty\"2\n\x05\x45rror\x12\x12\n\nerror_code\x18\x01 \x01(\x03\x12\x15\n\rerror_message\x18\x02 \x01(\t\";\n\x10SubscribeRequest\x12\'\n\rsubscriptions\x18\x01 \x03(\x0b\x32\x10.pb.Subscription\"]\n\x0cSubscription\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x1c\n\x07\x63hannel\x18\x02 \x01(\x0e\x32\x0b.pb.Channel\x12\x0e\n\x06symbol\x18\x03 \x01(\t\x12\r\n\x05\x64\x65pth\x18\x04 \x01(\x03\"\xa9\x02\n\x11SubscribeResponse\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\x1c\n\x07\x63hannel\x18\x03 \x01(\x0e\x32\x0b.pb.Channel\x12\x18\n\x05\x65vent\x18\x04 \x01(\x0e\x32\t.pb.Event\x12\x18\n\x05\x64\x65pth\x18\x05 \x01(\x0b\x32\t.pb.Depth\x12\x19\n\x06trades\x18\x06 \x03(\x0b\x32\t.pb.Trade\x12\x1a\n\x06ticker\x18\x07 \x01(\x0b\x32\n.pb.Ticker\x12\x19\n\x06orders\x18\x08 \x03(\x0b\x32\t.pb.Order\x12\x1d\n\x08\x62\x61lances\x18\t \x03(\x0b\x32\x0b.pb.Balance\x12\x15\n\rsubscribed_at\x18\n \x01(\x03\x12\x18\n\x05\x65rror\x18\x0b \x01(\x0b\x32\t.pb.Error\"g\n\x05\x44\x65pth\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\x1d\n\x04\x61sks\x18\x03 \x03(\x0b\x32\x0f.pb.PriceVolume\x12\x1d\n\x04\x62ids\x18\x04 \x03(\x0b\x32\x0f.pb.PriceVolume\",\n\x0bPriceVolume\x12\r\n\x05price\x18\x01 \x01(\x03\x12\x0e\n\x06volume\x18\x02 \x01(\x03\"\xc1\x01\n\x05Trade\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\t\x12\r\n\x05price\x18\x04 \x01(\x01\x12\x0e\n\x06volume\x18\x05 \x01(\x01\x12\x12\n\ncreated_at\x18\x06 \x01(\x03\x12\x16\n\x04side\x18\x07 \x01(\x0e\x32\x08.pb.Side\x12\x0b\n\x03\x66\x65\x65\x18\x08 \x01(\x01\x12\x14\n\x0c\x66\x65\x65_currency\x18\t \x01(\t\x12\r\n\x05maker\x18\n \x01(\x08\x12\r\n\x05trend\x18\x0b \x01(\t\"r\n\x06Ticker\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\x0c\n\x04open\x18\x03 \x01(\x01\x12\x0c\n\x04high\x18\x04 \x01(\x01\x12\x0b\n\x03low\x18\x05 \x01(\x01\x12\r\n\x05\x63lose\x18\x06 \x01(\x01\x12\x0e\n\x06volume\x18\x07 \x01(\x01\"\xb6\x02\n\x05Order\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\t\x12\x16\n\x04side\x18\x04 \x01(\x0e\x32\x08.pb.Side\x12!\n\norder_type\x18\x05 \x01(\x0e\x32\r.pb.OrderType\x12\r\n\x05price\x18\x06 \x01(\x01\x12\x12\n\nstop_price\x18\x07 \x01(\x01\x12\x11\n\tavg_price\x18\x08 \x01(\x01\x12\x0e\n\x06status\x18\t \x01(\t\x12\x12\n\ncreated_at\x18\n \x01(\x03\x12\x10\n\x08quantity\x18\x0b \x01(\x01\x12\x17\n\x0f\x65xecuted_volume\x18\x0c \x01(\x01\x12\x14\n\x0ctrades_count\x18\r \x01(\x03\x12\x17\n\x0f\x63lient_order_id\x18\x0e \x01(\t\x12\x10\n\x08group_id\x18\x0f \x01(\x03\"\xca\x01\n\x0bSubmitOrder\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\x16\n\x04side\x18\x03 \x01(\x0e\x32\x08.pb.Side\x12\x10\n\x08quantity\x18\x04 \x01(\x01\x12\r\n\x05price\x18\x05 \x01(\x01\x12\x12\n\nstop_price\x18\x06 \x01(\x01\x12!\n\norder_type\x18\x07 \x01(\x0e\x32\r.pb.OrderType\x12\x17\n\x0f\x63lient_order_id\x18\x08 \x01(\t\x12\x10\n\x08group_id\x18\t \x01(\x03\"P\n\x07\x42\x61lance\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x10\n\x08\x63urrency\x18\x02 \x01(\t\x12\x11\n\tavailable\x18\x03 \x01(\x01\x12\x0e\n\x06locked\x18\x04 \x01(\x01\";\n\x12SubmitOrderRequest\x12%\n\x0csubmit_order\x18\x01 \x01(\x0b\x32\x0f.pb.SubmitOrder\"I\n\x13SubmitOrderResponse\x12\x18\n\x05order\x18\x01 \x01(\x0b\x32\t.pb.Order\x12\x18\n\x05\x65rror\x18\x02 \x01(\x0b\x32\t.pb.Error\"K\n\x12\x43\x61ncelOrderRequest\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x17\n\x0f\x63lient_order_id\x18\x03 \x01(\t\"I\n\x13\x43\x61ncelOrderResponse\x12\x18\n\x05order\x18\x01 \x01(\x0b\x32\t.pb.Order\x12\x18\n\x05\x65rror\x18\x02 \x01(\x0b\x32\t.pb.Error\"J\n\x11QueryOrderRequest\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x17\n\x0f\x63lient_order_id\x18\x03 \x01(\t\"H\n\x12QueryOrderResponse\x12\x18\n\x05order\x18\x01 \x01(\x0b\x32\t.pb.Order\x12\x18\n\x05\x65rror\x18\x02 \x01(\x0b\x32\t.pb.Error\"\xaa\x01\n\x12QueryOrdersRequest\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\r\n\x05state\x18\x03 \x03(\t\x12\x10\n\x08order_by\x18\x04 \x01(\t\x12\x10\n\x08group_id\x18\x05 \x01(\x03\x12\x12\n\npagination\x18\x06 \x01(\x08\x12\x0c\n\x04page\x18\x07 \x01(\x03\x12\r\n\x05limit\x18\x08 \x01(\x03\x12\x0e\n\x06offset\x18\t \x01(\x03\"J\n\x13QueryOrdersResponse\x12\x19\n\x06orders\x18\x01 \x03(\x0b\x32\t.pb.Order\x12\x18\n\x05\x65rror\x18\x02 \x01(\x0b\x32\t.pb.Error\"\xb6\x01\n\x12QueryTradesRequest\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\x11\n\ttimestamp\x18\x03 \x01(\x03\x12\x0c\n\x04\x66rom\x18\x04 \x01(\x03\x12\n\n\x02to\x18\x05 \x01(\x03\x12\x10\n\x08order_by\x18\x06 \x01(\t\x12\x12\n\npagination\x18\x07 \x01(\x08\x12\x0c\n\x04page\x18\x08 \x01(\x03\x12\r\n\x05limit\x18\t \x01(\x03\x12\x0e\n\x06offset\x18\n \x01(\x03\"J\n\x13QueryTradesResponse\x12\x19\n\x06trades\x18\x01 \x03(\x0b\x32\t.pb.Trade\x12\x18\n\x05\x65rror\x18\x02 \x01(\x0b\x32\t.pb.Error\"j\n\x12QueryKLinesRequest\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\x10\n\x08interval\x18\x03 \x01(\t\x12\x11\n\ttimestamp\x18\x04 \x01(\x03\x12\r\n\x05limit\x18\x05 \x01(\x03\"J\n\x13QueryKLinesResponse\x12\x19\n\x06klines\x18\x01 \x03(\x0b\x32\t.pb.KLine\x12\x18\n\x05\x65rror\x18\x02 \x01(\x0b\x32\t.pb.Error\"\x9a\x01\n\x05KLine\x12\x10\n\x08\x65xchange\x18\x01 \x01(\t\x12\x0e\n\x06symbol\x18\x02 \x01(\t\x12\x11\n\ttimestamp\x18\x03 \x01(\x03\x12\x0c\n\x04open\x18\x04 \x01(\x01\x12\x0c\n\x04high\x18\x05 \x01(\x01\x12\x0b\n\x03low\x18\x06 \x01(\x01\x12\r\n\x05\x63lose\x18\x07 \x01(\x01\x12\x0e\n\x06volume\x18\x08 \x01(\x01\x12\x14\n\x0cquote_volume\x18\t \x01(\x01*\xe4\x01\n\x05\x45vent\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0e\n\nSUBSCRIBED\x10\x01\x12\x10\n\x0cUNSUBSCRIBED\x10\x02\x12\x0c\n\x08SNAPSHOT\x10\x03\x12\n\n\x06UPDATE\x10\x04\x12\x11\n\rAUTHENTICATED\x10\x05\x12\x12\n\x0eORDER_SNAPSHOT\x10\x06\x12\x10\n\x0cORDER_UPDATE\x10\x07\x12\x12\n\x0eTRADE_SNAPSHOT\x10\x08\x12\x10\n\x0cTRADE_UPDATE\x10\t\x12\x14\n\x10\x41\x43\x43OUNT_SNAPSHOT\x10\n\x12\x12\n\x0e\x41\x43\x43OUNT_UPDATE\x10\x0b\x12\t\n\x05\x45RROR\x10\x63*4\n\x07\x43hannel\x12\x08\n\x04\x42OOK\x10\x00\x12\t\n\x05TRADE\x10\x01\x12\n\n\x06TICKER\x10\x02\x12\x08\n\x04USER\x10\x03*\x19\n\x04Side\x12\x07\n\x03\x42UY\x10\x00\x12\x08\n\x04SELL\x10\x01*a\n\tOrderType\x12\n\n\x06MARKET\x10\x00\x12\t\n\x05LIMIT\x10\x01\x12\x0f\n\x0bSTOP_MARKET\x10\x02\x12\x0e\n\nSTOP_LIMIT\x10\x03\x12\r\n\tPOST_ONLY\x10\x04\x12\r\n\tIOC_LIMIT\x10\x05\x32\x93\x01\n\x11MarketDataService\x12<\n\tSubscribe\x12\x14.pb.SubscribeRequest\x1a\x15.pb.SubscribeResponse\"\x00\x30\x01\x12@\n\x0bQueryKLines\x12\x16.pb.QueryKLinesRequest\x1a\x17.pb.QueryKLinesResponse\"\x00\x32L\n\x0fUserDataService\x12\x39\n\x11SubscribeUserData\x12\t.pb.Empty\x1a\x15.pb.SubscribeResponse\"\x00\x30\x01\x32\xd7\x02\n\x0eTradingService\x12@\n\x0bSubmitOrder\x12\x16.pb.SubmitOrderRequest\x1a\x17.pb.SubmitOrderResponse\"\x00\x12@\n\x0b\x43\x61ncelOrder\x12\x16.pb.CancelOrderRequest\x1a\x17.pb.CancelOrderResponse\"\x00\x12=\n\nQueryOrder\x12\x15.pb.QueryOrderRequest\x1a\x16.pb.QueryOrderResponse\"\x00\x12@\n\x0bQueryOrders\x12\x16.pb.QueryOrdersRequest\x1a\x17.pb.QueryOrdersResponse\"\x00\x12@\n\x0bQueryTrades\x12\x16.pb.QueryTradesRequest\x1a\x17.pb.QueryTradesResponse\"\x00\x42\x07Z\x05../pbb\x06proto3')

_EVENT = DESCRIPTOR.enum_types_by_name['Event']
Event = enum_type_wrapper.EnumTypeWrapper(_EVENT)
_CHANNEL = DESCRIPTOR.enum_types_by_name['Channel']
Channel = enum_type_wrapper.EnumTypeWrapper(_CHANNEL)
_SIDE = DESCRIPTOR.enum_types_by_name['Side']
Side = enum_type_wrapper.EnumTypeWrapper(_SIDE)
_ORDERTYPE = DESCRIPTOR.enum_types_by_name['OrderType']
OrderType = enum_type_wrapper.EnumTypeWrapper(_ORDERTYPE)
UNKNOWN = 0
SUBSCRIBED = 1
UNSUBSCRIBED = 2
SNAPSHOT = 3
UPDATE = 4
AUTHENTICATED = 5
ORDER_SNAPSHOT = 6
ORDER_UPDATE = 7
TRADE_SNAPSHOT = 8
TRADE_UPDATE = 9
ACCOUNT_SNAPSHOT = 10
ACCOUNT_UPDATE = 11
ERROR = 99
BOOK = 0
TRADE = 1
TICKER = 2
USER = 3
BUY = 0
SELL = 1
MARKET = 0
LIMIT = 1
STOP_MARKET = 2
STOP_LIMIT = 3
POST_ONLY = 4
IOC_LIMIT = 5


_EMPTY = DESCRIPTOR.message_types_by_name['Empty']
_ERROR = DESCRIPTOR.message_types_by_name['Error']
_SUBSCRIBEREQUEST = DESCRIPTOR.message_types_by_name['SubscribeRequest']
_SUBSCRIPTION = DESCRIPTOR.message_types_by_name['Subscription']
_SUBSCRIBERESPONSE = DESCRIPTOR.message_types_by_name['SubscribeResponse']
_DEPTH = DESCRIPTOR.message_types_by_name['Depth']
_PRICEVOLUME = DESCRIPTOR.message_types_by_name['PriceVolume']
_TRADE = DESCRIPTOR.message_types_by_name['Trade']
_TICKER = DESCRIPTOR.message_types_by_name['Ticker']
_ORDER = DESCRIPTOR.message_types_by_name['Order']
_SUBMITORDER = DESCRIPTOR.message_types_by_name['SubmitOrder']
_BALANCE = DESCRIPTOR.message_types_by_name['Balance']
_SUBMITORDERREQUEST = DESCRIPTOR.message_types_by_name['SubmitOrderRequest']
_SUBMITORDERRESPONSE = DESCRIPTOR.message_types_by_name['SubmitOrderResponse']
_CANCELORDERREQUEST = DESCRIPTOR.message_types_by_name['CancelOrderRequest']
_CANCELORDERRESPONSE = DESCRIPTOR.message_types_by_name['CancelOrderResponse']
_QUERYORDERREQUEST = DESCRIPTOR.message_types_by_name['QueryOrderRequest']
_QUERYORDERRESPONSE = DESCRIPTOR.message_types_by_name['QueryOrderResponse']
_QUERYORDERSREQUEST = DESCRIPTOR.message_types_by_name['QueryOrdersRequest']
_QUERYORDERSRESPONSE = DESCRIPTOR.message_types_by_name['QueryOrdersResponse']
_QUERYTRADESREQUEST = DESCRIPTOR.message_types_by_name['QueryTradesRequest']
_QUERYTRADESRESPONSE = DESCRIPTOR.message_types_by_name['QueryTradesResponse']
_QUERYKLINESREQUEST = DESCRIPTOR.message_types_by_name['QueryKLinesRequest']
_QUERYKLINESRESPONSE = DESCRIPTOR.message_types_by_name['QueryKLinesResponse']
_KLINE = DESCRIPTOR.message_types_by_name['KLine']
Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), {
  'DESCRIPTOR' : _EMPTY,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Empty)
  })
_sym_db.RegisterMessage(Empty)

Error = _reflection.GeneratedProtocolMessageType('Error', (_message.Message,), {
  'DESCRIPTOR' : _ERROR,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Error)
  })
_sym_db.RegisterMessage(Error)

SubscribeRequest = _reflection.GeneratedProtocolMessageType('SubscribeRequest', (_message.Message,), {
  'DESCRIPTOR' : _SUBSCRIBEREQUEST,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.SubscribeRequest)
  })
_sym_db.RegisterMessage(SubscribeRequest)

Subscription = _reflection.GeneratedProtocolMessageType('Subscription', (_message.Message,), {
  'DESCRIPTOR' : _SUBSCRIPTION,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Subscription)
  })
_sym_db.RegisterMessage(Subscription)

SubscribeResponse = _reflection.GeneratedProtocolMessageType('SubscribeResponse', (_message.Message,), {
  'DESCRIPTOR' : _SUBSCRIBERESPONSE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.SubscribeResponse)
  })
_sym_db.RegisterMessage(SubscribeResponse)

Depth = _reflection.GeneratedProtocolMessageType('Depth', (_message.Message,), {
  'DESCRIPTOR' : _DEPTH,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Depth)
  })
_sym_db.RegisterMessage(Depth)

PriceVolume = _reflection.GeneratedProtocolMessageType('PriceVolume', (_message.Message,), {
  'DESCRIPTOR' : _PRICEVOLUME,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.PriceVolume)
  })
_sym_db.RegisterMessage(PriceVolume)

Trade = _reflection.GeneratedProtocolMessageType('Trade', (_message.Message,), {
  'DESCRIPTOR' : _TRADE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Trade)
  })
_sym_db.RegisterMessage(Trade)

Ticker = _reflection.GeneratedProtocolMessageType('Ticker', (_message.Message,), {
  'DESCRIPTOR' : _TICKER,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Ticker)
  })
_sym_db.RegisterMessage(Ticker)

Order = _reflection.GeneratedProtocolMessageType('Order', (_message.Message,), {
  'DESCRIPTOR' : _ORDER,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Order)
  })
_sym_db.RegisterMessage(Order)

SubmitOrder = _reflection.GeneratedProtocolMessageType('SubmitOrder', (_message.Message,), {
  'DESCRIPTOR' : _SUBMITORDER,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.SubmitOrder)
  })
_sym_db.RegisterMessage(SubmitOrder)

Balance = _reflection.GeneratedProtocolMessageType('Balance', (_message.Message,), {
  'DESCRIPTOR' : _BALANCE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.Balance)
  })
_sym_db.RegisterMessage(Balance)

SubmitOrderRequest = _reflection.GeneratedProtocolMessageType('SubmitOrderRequest', (_message.Message,), {
  'DESCRIPTOR' : _SUBMITORDERREQUEST,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.SubmitOrderRequest)
  })
_sym_db.RegisterMessage(SubmitOrderRequest)

SubmitOrderResponse = _reflection.GeneratedProtocolMessageType('SubmitOrderResponse', (_message.Message,), {
  'DESCRIPTOR' : _SUBMITORDERRESPONSE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.SubmitOrderResponse)
  })
_sym_db.RegisterMessage(SubmitOrderResponse)

CancelOrderRequest = _reflection.GeneratedProtocolMessageType('CancelOrderRequest', (_message.Message,), {
  'DESCRIPTOR' : _CANCELORDERREQUEST,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.CancelOrderRequest)
  })
_sym_db.RegisterMessage(CancelOrderRequest)

CancelOrderResponse = _reflection.GeneratedProtocolMessageType('CancelOrderResponse', (_message.Message,), {
  'DESCRIPTOR' : _CANCELORDERRESPONSE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.CancelOrderResponse)
  })
_sym_db.RegisterMessage(CancelOrderResponse)

QueryOrderRequest = _reflection.GeneratedProtocolMessageType('QueryOrderRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUERYORDERREQUEST,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryOrderRequest)
  })
_sym_db.RegisterMessage(QueryOrderRequest)

QueryOrderResponse = _reflection.GeneratedProtocolMessageType('QueryOrderResponse', (_message.Message,), {
  'DESCRIPTOR' : _QUERYORDERRESPONSE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryOrderResponse)
  })
_sym_db.RegisterMessage(QueryOrderResponse)

QueryOrdersRequest = _reflection.GeneratedProtocolMessageType('QueryOrdersRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUERYORDERSREQUEST,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryOrdersRequest)
  })
_sym_db.RegisterMessage(QueryOrdersRequest)

QueryOrdersResponse = _reflection.GeneratedProtocolMessageType('QueryOrdersResponse', (_message.Message,), {
  'DESCRIPTOR' : _QUERYORDERSRESPONSE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryOrdersResponse)
  })
_sym_db.RegisterMessage(QueryOrdersResponse)

QueryTradesRequest = _reflection.GeneratedProtocolMessageType('QueryTradesRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUERYTRADESREQUEST,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryTradesRequest)
  })
_sym_db.RegisterMessage(QueryTradesRequest)

QueryTradesResponse = _reflection.GeneratedProtocolMessageType('QueryTradesResponse', (_message.Message,), {
  'DESCRIPTOR' : _QUERYTRADESRESPONSE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryTradesResponse)
  })
_sym_db.RegisterMessage(QueryTradesResponse)

QueryKLinesRequest = _reflection.GeneratedProtocolMessageType('QueryKLinesRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUERYKLINESREQUEST,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryKLinesRequest)
  })
_sym_db.RegisterMessage(QueryKLinesRequest)

QueryKLinesResponse = _reflection.GeneratedProtocolMessageType('QueryKLinesResponse', (_message.Message,), {
  'DESCRIPTOR' : _QUERYKLINESRESPONSE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.QueryKLinesResponse)
  })
_sym_db.RegisterMessage(QueryKLinesResponse)

KLine = _reflection.GeneratedProtocolMessageType('KLine', (_message.Message,), {
  'DESCRIPTOR' : _KLINE,
  '__module__' : 'bbgo_pb2'
  # @@protoc_insertion_point(class_scope:pb.KLine)
  })
_sym_db.RegisterMessage(KLine)

_MARKETDATASERVICE = DESCRIPTOR.services_by_name['MarketDataService']
_USERDATASERVICE = DESCRIPTOR.services_by_name['UserDataService']
_TRADINGSERVICE = DESCRIPTOR.services_by_name['TradingService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\005../pb'
  _EVENT._serialized_start=2888
  _EVENT._serialized_end=3116
  _CHANNEL._serialized_start=3118
  _CHANNEL._serialized_end=3170
  _SIDE._serialized_start=3172
  _SIDE._serialized_end=3197
  _ORDERTYPE._serialized_start=3199
  _ORDERTYPE._serialized_end=3296
  _EMPTY._serialized_start=18
  _EMPTY._serialized_end=25
  _ERROR._serialized_start=27
  _ERROR._serialized_end=77
  _SUBSCRIBEREQUEST._serialized_start=79
  _SUBSCRIBEREQUEST._serialized_end=138
  _SUBSCRIPTION._serialized_start=140
  _SUBSCRIPTION._serialized_end=233
  _SUBSCRIBERESPONSE._serialized_start=236
  _SUBSCRIBERESPONSE._serialized_end=533
  _DEPTH._serialized_start=535
  _DEPTH._serialized_end=638
  _PRICEVOLUME._serialized_start=640
  _PRICEVOLUME._serialized_end=684
  _TRADE._serialized_start=687
  _TRADE._serialized_end=880
  _TICKER._serialized_start=882
  _TICKER._serialized_end=996
  _ORDER._serialized_start=999
  _ORDER._serialized_end=1309
  _SUBMITORDER._serialized_start=1312
  _SUBMITORDER._serialized_end=1514
  _BALANCE._serialized_start=1516
  _BALANCE._serialized_end=1596
  _SUBMITORDERREQUEST._serialized_start=1598
  _SUBMITORDERREQUEST._serialized_end=1657
  _SUBMITORDERRESPONSE._serialized_start=1659
  _SUBMITORDERRESPONSE._serialized_end=1732
  _CANCELORDERREQUEST._serialized_start=1734
  _CANCELORDERREQUEST._serialized_end=1809
  _CANCELORDERRESPONSE._serialized_start=1811
  _CANCELORDERRESPONSE._serialized_end=1884
  _QUERYORDERREQUEST._serialized_start=1886
  _QUERYORDERREQUEST._serialized_end=1960
  _QUERYORDERRESPONSE._serialized_start=1962
  _QUERYORDERRESPONSE._serialized_end=2034
  _QUERYORDERSREQUEST._serialized_start=2037
  _QUERYORDERSREQUEST._serialized_end=2207
  _QUERYORDERSRESPONSE._serialized_start=2209
  _QUERYORDERSRESPONSE._serialized_end=2283
  _QUERYTRADESREQUEST._serialized_start=2286
  _QUERYTRADESREQUEST._serialized_end=2468
  _QUERYTRADESRESPONSE._serialized_start=2470
  _QUERYTRADESRESPONSE._serialized_end=2544
  _QUERYKLINESREQUEST._serialized_start=2546
  _QUERYKLINESREQUEST._serialized_end=2652
  _QUERYKLINESRESPONSE._serialized_start=2654
  _QUERYKLINESRESPONSE._serialized_end=2728
  _KLINE._serialized_start=2731
  _KLINE._serialized_end=2885
  _MARKETDATASERVICE._serialized_start=3299
  _MARKETDATASERVICE._serialized_end=3446
  _USERDATASERVICE._serialized_start=3448
  _USERDATASERVICE._serialized_end=3524
  _TRADINGSERVICE._serialized_start=3527
  _TRADINGSERVICE._serialized_end=3870
# @@protoc_insertion_point(module_scope)
