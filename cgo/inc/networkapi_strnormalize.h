//
//  networkapi_strnormalize.h
//  BLNetwork
//
//  Created by yzm157 on 14-4-23.
//  Copyright (c) 2014年 BroadLink. All rights reserved.
//

#ifndef BLNetwork_networkapi_strnormalize_h
#define BLNetwork_networkapi_strnormalize_h

#ifdef __cplusplus
extern "C" {
#endif
    
#define NETWORKAPI_SNO_TO_LOWER        1
#define NETWORKAPI_SNO_TO_UPPER        2
#define NETWORKAPI_SNO_TO_HALF         4
#define NETWORKAPI_SNO_TO_SIMPLIFIED   8
    
void networkapi_str_normalize_init();
void networkapi_str_normalize_gbk(char *text, unsigned options);
void networkapi_str_normalize_utf8(char *text, unsigned options);
    
int networkapi_gbk_to_utf8(const char *from, unsigned int from_len, char **to, unsigned int *to_len);
int networkapi_utf8_to_gbk(const char *from, unsigned int from_len, char **to, unsigned int *to_len);
    
#ifdef __cplusplus
}
#endif

#endif
