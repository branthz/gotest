//
//  networkapi_define.h
//  NetworkAPI
//
//  Created by yzm157 on 14/12/3.
//  Copyright (c) 2014年 BroadLink. All rights reserved.
//

#ifndef NetworkAPI_networkapi_define_h
#define NetworkAPI_networkapi_define_h

#ifdef __cplusplus
extern "C"
{
#endif

//若为发布版本，则设置为非0值
#define NETWORKAPI_RELEASE                      1
    
//若为局域网版本，则无法通过远程使用，且无需使用license初始化
//#define NETWORKAPI_LOCALCONTROL_VERSION

/*eControl使用时为1(无需使用license),其他位0(需要userlicense以及typelicense)*/
//#define NETWORKAPI_BROADLINK_SUPER_USER
//#define NETWORKAPI_FOR_SERVER
#define NETWORKAPI_VERSION                  "2.0.0"
#define NETWORKAPI_BUILD                    "201502111420-prealpha"

#ifndef NETWORKAPI_BROADLINK_SUPER_USER
#ifndef NETWORKAPI_LOCALCONTROL_VERSION
#define NETWORKAPI_TOKEN_BUCKET_INTERVEL    10
#define NETWORKAPI_TOKEN_BUCKET_COUNT       20
#endif
#endif
    
#ifdef __cplusplus
}
#endif

#endif
