//
//  dnanl.h
//
//  Copyright (c) 2015年 BroadLink. All rights reserved.
//

#ifndef __DNANL_H__
#define __DNANL_H__

#ifdef __cplusplus
extern "C"
{
#endif

/*
 * !!!! All string returned must be released by free function,!!!
 * !!!! because they are malloced in the sdk                  !!!
 * !!!! All APIs are block mode                               !!!
 */

/**
  * SDKInit is used to do some bootstrap thing for the whole SDK.
  *
  */
extern char *SDKInit(const char *jsonstr);

/**
  * let device get WLAN config.
  * 
  */
extern char *deviceEasyConfig(const char *jsonstr);

/**
  * you can stop the easyconfig process.
  */
extern char *deviceEasyConfigCancel();

/**
  * do device discovery work on the same LAN.
  */
extern char *deviceProbe(const char *descstr);

/**
  * you must pair with the device before you can control it.
  * device will give you a key to access it.
  */
extern char *devicePair(const char *cachestr, int retrycount);

/**
  * to see if device is on server, you only need this api 
  * if you want to control device remotely.
  */
extern char *deviceOnServer(const char *cachestr, int retrycount);

/**
  * if you have more then one of the same type of device.
  * you shoud need this api to see if devices are on server.
  */
extern char *deviceListOnServer(const char *macliststr, unsigned short type, int retrycount);

/**
  * change device name and lock state.
  */
extern char *deviceConfig(const char *cachestr, const char *information);

/**
  * get the time of the server connected with the device.
  * you only need use this api if you want to do some timer work.
  */
extern char *deviceServerTime(const char *cachestr, int retrycount);

/**
  * for now, just return firmware version.
  */
extern char *deviceFirmwareInformation(const char *cachestr, int retrycount);

/**
  * upgrade the firmware of a device.
  */
extern char *deviceFirmwareUpgrade(const char *cachestr, const char *url, int retrycount);

/**
  * device control
  */
extern char *dnaControl(const char *cachestr, const char *string);
    
/**
  * Get dna device's IFTTT data
  */
extern char *dnaIFTTTData(const char *cachestr, const char *string);
    
#ifdef __cplusplus
}
#endif

#endif /* defined(__DNANL_H__) */
