//
//  networkapi_common.h
//  NetworkAPI
//
//  Created by yzm157 on 14/10/29.
//  Copyright (c) 2014年 BroadLink. All rights reserved.
//

#ifndef __NetworkAPI__networkapi_common__
#define __NetworkAPI__networkapi_common__

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <pthread.h>
#include <string.h>
#include <arpa/inet.h>
#include "networkapi_define.h"
#include "networkapi_json.h"
#include "networkapi_token.h"

#ifdef __cplusplus
extern "C"
{
#endif

typedef int8_t      INT8;
typedef int16_t     INT16;
typedef int32_t     INT32;
typedef int64_t     INT64;
typedef uint8_t     UINT8;
typedef uint16_t    UINT16;
typedef uint32_t    UINT32;
typedef uint64_t    UINT64;

#if (NETWORKAPI_RELEASE == 1)
#define NETWORKAPI_TEST_LOG_ERROR               0
#define NETWORKAPI_TEST_LOG_DEBUG               0
#define NETWORKAPI_TEST_LOG_WARN                0
#else
#define NETWORKAPI_TEST_LOG_ERROR               1
#define NETWORKAPI_TEST_LOG_DEBUG               1
#define NETWORKAPI_TEST_LOG_WARN                1
#endif

#ifndef __LOCAL_COMPILE__
#else
#define loge(format, ...)  do {if (NETWORKAPI_TEST_LOG_ERROR) {printf("[Error]: %s, %s, %d\r\n" format "\r\n", __FILE__, __FUNCTION__, __LINE__, ##__VA_ARGS__);fflush(stdout);}}while(0)
#define logd(format, ...)  do {if (NETWORKAPI_TEST_LOG_DEBUG) {printf("[Debug]: %s, %s, %d\r\n" format "\r\n", __FILE__, __FUNCTION__, __LINE__, ##__VA_ARGS__);fflush(stdout);}}while(0)
#define logw(format, ...)  do {if (NETWORKAPI_TEST_LOG_WARN) {printf("[Warning]: %s, %s, %d\r\n" format "\r\n", __FILE__, __FUNCTION__, __LINE__ , ##__VA_ARGS__);fflush(stdout);}}while(0)
#endif

#define OLD_MODULE_TYPE_END     10000   //一代设备的结束值

typedef enum
{
    BROADLINK_TYPE_SP1 = 0,             /*Broadlink Co., Ltd.   SP1*/
    BROADLINK_TYPE_RM1 = 10000,         /*Broadlink Co., Ltd.   RM1*/
    BROADLINK_TYPE_SP2 = 10001,         /*Broadlink Co., Ltd.   SP2*/
    BROADLINK_TYPE_RMPRO = 10002,       /*Broadlink Co., Ltd.   RMPro*/
    BROADLINK_TYPE_X1 = 10003,          /*Broadlink Co., Ltd.   X1*/
    BROADLINK_TYPE_A1 = 10004,          /*Broadlink Co., Ltd.   A1*/
    BROADLINK_TYPE_SPMINI_OLD = 10016,  /*Broadlink Co., Ltd.   SPMini_Old*/
    BROADLINK_TYPE_SPMINI = 10024,      /*Broadlink Co., Ltd.   SPMini*/
    BROADLINK_TYPE_SPOEM = 10032,       /*Broadlink Co., Ltd.   SPOEM*/
    
    TCL_TYPE_T3_V1 = 7,                 /*Foshan TCL Household Appliances(Nanhai) Co.,Ltd.  T3_V1*/
    TCL_TYPE_T3_V2 = 20007,             /*Foshan TCL Household Appliances(Nanhai) Co.,Ltd.  T3_V2*/
    HONYAR_TYPE_MS3 = 10019,            /*Honyar Intelligent technology Co.,Ltd. MS3*/
    HONYAR_TYPE_IHLD800 = 10020,        /*Honyar Intelligent technology Co.,Ltd. IHLD800*/
    HONYAR_TYPE_IHLD480 = 10021,        /*Honyar Intelligent technology Co.,Ltd. IHLD480*/
    HONYAR_TYPE_IHL1301 = 10022,        /*Honyar Intelligent technology Co.,Ltd. IHL1301*/
    HONYAR_TYPE_IHL1302 = 10023,        /*Honyar Intelligent technology Co.,Ltd. IHL1302*/
}device_type_e;

/*模块的名称和锁定状态*/
typedef struct device_info_t
{
    UINT8 name[63];
    UINT8 lock;
}__attribute__((packed))device_info_t;

/*设备配对返回结构体*/
typedef struct pair_info_t
{
    INT32 control;
    UINT8 key[16];
}__attribute__((packed))pair_info_t;

/*probe返回至APP的结构体*/
typedef struct probe_cache_t
{
    UINT8 mac[6];
    UINT16 device_type;
    UINT16 sub_device;
    UINT32 control_pwd;
    UINT64 magic;
    device_info_t info;
    pair_info_t pair;
    struct sockaddr_in from;
    UINT8 is_lan;                       //是否本地的标记
}__attribute__((packed))probe_cache_t;
    
/*batch onserver check*/
typedef struct batch_onserver_check_t
{
    UINT8 mac[6];
    INT16 status;
}__attribute__((packed))batch_onserver_check_t;

#define MAX_DEVICETYPE_SUPPORT_CNT  50
typedef struct sdk_user_info_t
{
    UINT8 license[256];             //用户license;
    UINT8 illegal;                  //license的有效性
    UINT8 super_user;               //超级用户，没有设备类型限制
    UINT8 typecount;
    UINT16 type[MAX_DEVICETYPE_SUPPORT_CNT];
    time_t activation_time;         //license认证时间
}__attribute__((packed))sdk_user_info_t;

typedef struct auth_info_t
{
    UINT8 terminal_name[24];
    UINT8 terminal_udid[24];
    UINT16 terminal_type;
}__attribute__((packed))auth_info_t;

//全局变量
extern auth_info_t globalauth;
#ifndef NETWORKAPI_BROADLINK_SUPER_USER
#ifndef NETWORKAPI_LOCALCONTROL_VERSION
extern struct networkapi_token_bucket_t bucket;
#endif
extern sdk_user_info_t globaluser;               //sdk用户信息
#endif

/*判断大小端*/
static inline int isbigendian()
{
    int retval = 0;
    union endian_u
    {
        char a;
        UINT16 b;
    }endian;
    
    endian.b = 1;
    
    if (endian.a == 0)
        retval = 1;
    
    return retval;
}

#define bl_swab16(x)    ((UINT16)((((UINT16)(x) & 0xff00) >> 8) | (((UINT16)(x) & 0x00ff) << 8)))
#define bl_swab32(x)    ((UINT32)((((UINT32)(x) & 0xff000000) >> 24) | (((UINT32)(x) & 0x00ff0000) >>  8) | (((UINT32)(x) & 0x0000ff00) <<  8) | (((UINT32)(x) & 0x000000ff) << 24)))
#define bl_swab64(x)    ((UINT64)((((UINT64)(x) & 0xff00000000000000ULL) >> 56) | \
                        (((UINT64)(x) & 0x00ff000000000000ULL) >> 40) | \
                        (((UINT64)(x) & 0x0000ff0000000000ULL) >> 24) | \
                        (((UINT64)(x) & 0x000000ff00000000ULL) >>  8) | \
                        (((UINT64)(x) & 0x00000000ff000000ULL) <<  8) | \
                        (((UINT64)(x) & 0x0000000000ff0000ULL) << 24) | \
                        (((UINT64)(x) & 0x000000000000ff00ULL) << 40) | \
                        (((UINT64)(x) & 0x00000000000000ffULL) << 56)))

/*short: little endian to big endian*/
#define _l2b16(x)    isbigendian() ? bl_swab16((x)) : (x)
#define _l2b32(x)    isbigendian() ? bl_swab32((x)) : (x)
#define _l2b64(x)    isbigendian() ? bl_swab64((x)) : (x)
    
#ifdef __cplusplus
}
#endif

#endif /* defined(__NetworkAPI__BLCommon__) */
