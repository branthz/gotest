//
//  networkapi_errno.h
//  NetworkAPI
//
//  Created by yzm157 on 14/10/30.
//  Copyright (c) 2014年 BroadLink. All rights reserved.
//

#ifndef NetworkAPI_networkapi_errno_h
#define NetworkAPI_networkapi_errno_h

#ifdef __cplusplus
extern "C"
{
#endif

typedef enum broadlink_errno_e
{
    BL_IFTTT_DEVTYPE_ERR    = -1027,                /*不支持IFTTT的设备*/
    BL_DATA_ERR             = -1026,                /*传输数据过长*/
    BL_NOT_SUPPORT_REMOTE   = -1025,                /*SDK版本不支持远程操作*/
    BL_OPERATING_FAST       = -1024,                /*操作过于频繁*/
    BL_LICENSE_REJECT       = -1023,                /*无效license*/
    BL_FUNCTION_FAIL        = -1022,                /*不支持的param*/
    BL_USER_MSG_FAIL        = -1021,                /*用户类型错误*/
    BL_FILE_FAIL            = -1020,                /*bl文件有误*/
    BL_PATTERN_FAIL         = -1019,                /*pat文件有误*/
    BL_MALLOC_FAIL          = -1018,                /*分配内存失败*/
    BL_INFO_ERR             = -1017,                /*输入的设备信息有误*/
    BL_JSON_TYPE_ERR        = -1016,                /*JSON字符串的数据类型有误*/
    BL_JSON_ERR             = -1015,                /*传入的JSON字符串有误*/
    BL_INIT_FAIL            = -1014,                /*网络库初始化失败*/
    BL_DNS_PARSE_FAIL       = -1013,                /*域名解析失败*/
    BL_CONTROL_ID_FAIL      = -1012,                /*设备控制ID错误，设备已经复位且控制终端未在局域网与设备配对配对*/
    BL_AES_CHECK_FAIL       = -1011,                /*接收数据解密校验失败*/
    BL_AES_LEN_FAIL         = -1010,                /*接收数据解密长度失败*/
    BL_HEAD_MSG_FAIL        = -1009,                /*网络消息类型错误*/
    BL_HEAD_CHECK_FAIL      = -1008,                /*接收数据包校验失败*/
    BL_RECV_LEN_FAIL        = -1007,                /*接收数据包长度有误*/
    BL_SELECT_FAIL          = -1006,                /*socket操作失败*/
    BL_SOCKET_SEND_FAIL     = -1005,                /*socket发包失败*/
    BL_SET_SOCKET_OPT_FAIL  = -1004,                /*设置socket属性失败*/
    BL_CREATE_SOCKET_FAIL   = -1003,                /*创建socket失败*/
    BL_EASYCONFIG_CANCEL    = -1002,                /*取消easyconfig*/
    BL_NOT_LAN              = -1001,                /*设备不在局域网*/
    BL_TIMEOUT              = -1000,                /*超时*/
    BL_SUCCESS              = 0,
}broadlink_errno_e;
    
#ifdef __cplusplus
}
#endif

#endif
