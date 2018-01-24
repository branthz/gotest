//
//  networkapi_command.h
//  NetworkAPI
//
//  Created by yzm157 on 14/11/29.
//  Copyright (c) 2014年 BroadLink. All rights reserved.
//

#ifndef NetworkAPI_networkapi_command_h
#define NetworkAPI_networkapi_command_h

#include "networkapi_common.h"
#include "networkapi_errno.h"
#include "networkapi_strnormalize.h"

#ifdef __cplusplus
extern "C"
{
#endif

#define NETWORKAPI_FOR_SERVER

#define NETWORKAPI_JSON_RESULT(JSON, errcode, message)   {\
BLJSON_AddNumberToObject((JSON), "code", (errcode));\
BLJSON_AddStringToObject((JSON), "msg", (message));\
goto end;\
}

/*
#define NETWORKAPI_JSON_ADD_DEVINFO(json, name, lock)  {\
char temp[64];\
BLJSON *item = BLJSON_CreateObject();\
snprintf(temp, sizeof(temp), "%02x:%02x:%02x:%02x:%02x:%02x", cache->mac[5], cache->mac[4], cache->mac[3], cache->mac[2], cache->mac[1], cache->mac[0]);\
BLJSON_AddStringToObject(item, "mac", temp);\
BLJSON_AddStringToObject(item, "name", (const char *)(name));\
BLJSON_AddBoolToObject(item, "lock", (lock));\
BLJSON_AddNumberToObject(item, "type", cache->device_type);\
BLJSON_AddNumberToObject(item, "password", cache->control_pwd);\
BLJSON_AddNumberToObject(item, "id", cache->pair.control);\
snprintf(temp, sizeof(temp), "%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", cache->pair.key[0], cache->pair.key[1], cache->pair.key[2], cache->pair.key[3], cache->pair.key[4], cache->pair.key[5], cache->pair.key[6], cache->pair.key[7], cache->pair.key[8], cache->pair.key[9], cache->pair.key[10], cache->pair.key[11], cache->pair.key[12], cache->pair.key[13], cache->pair.key[14], cache->pair.key[15]);\
BLJSON_AddStringToObject(item, "key", temp);\
if (cache->device_type <= OLD_MODULE_TYPE_END)\
BLJSON_AddNumberToObject(item, "magiccode", cache->magic);\
if (cache->is_lan)\
{\
snprintf(temp, sizeof(temp), "%s:%d", inet_ntoa(cache->from.sin_addr), ntohs(cache->from.sin_port));\
BLJSON_AddStringToObject(item, "lanaddr", temp);\
}\
BLJSON_AddItemToObject((json), "devinfo", item);\
}

#define NETWORKAPI_JSON_ADD_DEVINFO2(json, name, lock)  {\
char temp[64];\
BLJSON *item = BLJSON_CreateObject();\
snprintf(temp, sizeof(temp), "%02x:%02x:%02x:%02x:%02x:%02x", cache.mac[5], cache.mac[4], cache.mac[3], cache.mac[2], cache.mac[1], cache.mac[0]);\
BLJSON_AddStringToObject(item, "mac", temp);\
BLJSON_AddStringToObject(item, "name", (const char *)(name));\
BLJSON_AddBoolToObject(item, "lock", (lock));\
BLJSON_AddNumberToObject(item, "type", cache.device_type);\
BLJSON_AddNumberToObject(item, "password", cache.control_pwd);\
BLJSON_AddNumberToObject(item, "id", cache.pair.control);\
snprintf(temp, sizeof(temp), "%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", cache.pair.key[0], cache.pair.key[1], cache.pair.key[2], cache.pair.key[3], cache.pair.key[4], cache.pair.key[5], cache.pair.key[6], cache.pair.key[7], cache.pair.key[8], cache.pair.key[9], cache.pair.key[10], cache.pair.key[11], cache.pair.key[12], cache.pair.key[13], cache.pair.key[14], cache.pair.key[15]);\
BLJSON_AddStringToObject(item, "key", temp);\
if (cache.device_type <= OLD_MODULE_TYPE_END)\
BLJSON_AddNumberToObject(item, "magiccode", cache.magic);\
if (cache.is_lan)\
{\
snprintf(temp, sizeof(temp), "%s:%d", inet_ntoa(cache.from.sin_addr), ntohs(cache.from.sin_port));\
BLJSON_AddStringToObject(item, "lanaddr", temp);\
}\
BLJSON_AddItemToObject((json), "devinfo", item);\
}
 */

static inline int encoding_utf2gbk(char *buf, int buf_len)
{
    UINT32 utf8_len = (UINT32)strlen(buf);
    UINT32 gbkbuffer_len = utf8_len * 2 + 1;
    char *gbkbuffer = (char *)malloc(gbkbuffer_len);
    if (gbkbuffer == NULL)
        return BL_MALLOC_FAIL;
    memset(gbkbuffer, 0, gbkbuffer_len);
    networkapi_utf8_to_gbk(buf, utf8_len, &gbkbuffer, &gbkbuffer_len);
    snprintf(buf, buf_len, "%s", gbkbuffer);
    free(gbkbuffer);
    
    return 0;
}

static inline int encoding_gbk2utf(char *buf, int buf_len)
{
    UINT32 gbk_len = (UINT32)strlen(buf);
    UINT32 newlen = gbk_len * 3 + 1;
    char *utf8buffer = (char *)malloc(newlen);
    if (utf8buffer == NULL)
        return BL_MALLOC_FAIL;
    memset(utf8buffer, 0, newlen);
    networkapi_gbk_to_utf8(buf, gbk_len, &utf8buffer, &newlen);
    snprintf(buf, buf_len, "%s", utf8buffer);
    free(utf8buffer);
    
    return 0;
}

/*mac string to binary*/
static inline int mac2binary(const char *string, UINT8 mac[6])
{
    int temp[6];
    int i;
    
    if (NULL == string)
        return BL_INFO_ERR;
    
    if (sscanf(string, "%02x:%02x:%02x:%02x:%02x:%02x"\
               , &temp[5], &temp[4], &temp[3]\
               , &temp[2], &temp[1], &temp[0]) < 6)
        return BL_INFO_ERR;
    
    for (i=0; i<6; i++)
        mac[i] = temp[i] & 0xff;
    
    return 0;
}

/*info string to probe_cache_t*/
static inline int string2cache(const char *jsonstr, probe_cache_t *cache)
{
    BLJSON *parse = NULL;
    BLJSON *item = NULL;
    int i=0;
    int key[16];
    int ip[4];
    int port;
    char temp[20];
    
    if (NULL == jsonstr || NULL == cache)
        return BL_INFO_ERR;
    
    if (NULL == (parse = BLJSON_Parse(jsonstr)))
        return BL_INFO_ERR;
    
    memset(cache, 0, sizeof(probe_cache_t));
    
    if (NULL == (item = BLJSON_GetObjectItem(parse, "mac")))
        goto err;
    if (item->type != BLJSON_String)
        goto err;
    if (mac2binary(item->valuestring, cache->mac) < 0)
        goto err;
    
    if (NULL != (item = BLJSON_GetObjectItem(parse, "name")))
    {
        if (item->type != BLJSON_String)
            goto err;
        snprintf((char *)cache->info.name, sizeof(cache->info.name), "%s", item->valuestring);
    }
    
    if (NULL != (item = BLJSON_GetObjectItem(parse, "lock")))
    {
        if (item->type != BLJSON_True && item->type != BLJSON_False)
            goto err;
        cache->info.lock = (item->type == BLJSON_True) ? 1 : 0;
    }
    
    if (NULL == (item = BLJSON_GetObjectItem(parse, "type")))
        goto err;
    if (item->type != BLJSON_Number)
        goto err;
    cache->device_type = (UINT16)item->valueint;
    
    if (NULL != (item = BLJSON_GetObjectItem(parse, "subdevice")))
    {
        if (item->type == BLJSON_Number)
            cache->sub_device = (UINT16)item->valueint;
    }
    
    if (NULL == (item = BLJSON_GetObjectItem(parse, "password")))
        goto err;
    if (item->type != BLJSON_Number)
        goto err;
    cache->control_pwd = (UINT32)item->valueint;

    if (NULL != (item = BLJSON_GetObjectItem(parse, "magiccode")))
    {
        if (item->type != BLJSON_String)
            goto err;
        cache->magic = (UINT64)atoll(item->valuestring);
    }

    if (NULL != (item = BLJSON_GetObjectItem(parse, "id")))
    {
        if (item->type != BLJSON_Number)
            goto err;
        cache->pair.control = (INT32)item->valueint;
    }
    
    if (NULL != (item = BLJSON_GetObjectItem(parse, "key")))
    {
        if (item->type != BLJSON_String)
            goto err;
        if (sscanf(item->valuestring, "%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", &key[0], &key[1], &key[2], &key[3], &key[4], &key[5], &key[6], &key[7], &key[8], &key[9], &key[10], &key[11], &key[12], &key[13], &key[14], &key[15]) < 16)
            goto err;
        for (i=0; i<16; i++)
            cache->pair.key[i] = key[i] & 0xff;
    }
    
    if (NULL != (item = BLJSON_GetObjectItem(parse, "lanaddr")))
    {
        if (item->type != BLJSON_String)
            goto err;
        if (sscanf(item->valuestring, "%d.%d.%d.%d:%d", &ip[0], &ip[1], &ip[2], &ip[3], &port) < 5)
        {
            goto err;
        }
        snprintf(temp, sizeof(temp), "%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3]);
        cache->from.sin_family = AF_INET;
        cache->from.sin_addr.s_addr = inet_addr(temp);
        cache->from.sin_port = htons(port);
        cache->is_lan = 1;
    }
    
    BLJSON_Delete(parse);
    return 0;
    
err:
    BLJSON_Delete(parse);
    return BL_INFO_ERR;
}

extern char *networkapi_init(const char *license, auth_info_t *info);
extern char *networkapi_device_easyconfig(const char *string);
extern char *networkapi_device_easyconfig_cancel();
extern char *networkapi_device_probe(const char *descstr);
extern char *networkapi_device_pair(const char *cachestr, int count);
extern char *networkapi_device_state(const char *cachestr, int count);
extern char *networkapi_batch_onserver(unsigned short type, const char *macstr, int count);
extern char *networkapi_device_set_info(const char *cachestr, const char *information);
extern char *networkapi_device_server_time(const char *cachestr, int count);
extern char *networkapi_device_firmware_version(const char *cachestr, unsigned int local, unsigned int remote, int count);
extern char *networkapi_device_firmware_upgrade(const char *cachestr, const char *url, unsigned int local, unsigned int remote, int count);
extern char *networkapi_device_control(const char *cachestr, const char *string);
extern char *networkapi_ifttt_data(const char *cachestr, const char *string);
//extern char *networkapi_device_send(const char *cachestr, const char *data, unsigned short type, unsigned int local, unsigned int remote, int count);
    
    
/*only for server*/
#ifdef NETWORKAPI_FOR_SERVER
extern int cloudserv_json2c(const char *path, const char *cachestr, const char *string, INT32 *cmd, UINT8 buf[1460], UINT32 *datalen);
extern char *cloudserv_c2json(const char *path, const char *cachestr, UINT8 buf[1460], UINT32 datalen, INT32 cmd);
#endif

extern char *networkapi_not_support(const char *param);
    
#ifdef __cplusplus
}
#endif

#endif
