/*
  Copyright (c) 2009 Dave Gamble
 
  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:
 
  The above copyright notice and this permission notice shall be included in
  all copies or substantial portions of the Software.
 
  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
  THE SOFTWARE.
*/

#ifndef BLJSON__h
#define BLJSON__h

#ifdef __cplusplus
extern "C"
{
#endif

/* BLJSON Types: */
#define BLJSON_False 0
#define BLJSON_True 1
#define BLJSON_NULL 2
#define BLJSON_Number 3
#define BLJSON_String 4
#define BLJSON_Array 5
#define BLJSON_Object 6
	
#define BLJSON_IsReference 256

/* The BLJSON structure: */
typedef struct BLJSON {
	struct BLJSON *next,*prev;	/* next/prev allow you to walk array/object chains. Alternatively, use GetArraySize/GetArrayItem/GetObjectItem */
	struct BLJSON *child;		/* An array or object item will have a child pointer pointing to a chain of the items in the array/object. */

	int type;					/* The type of the item, as above. */

	char *valuestring;			/* The item's string, if type==BLJSON_String */
	int valueint;				/* The item's number, if type==BLJSON_Number */
	double valuedouble;			/* The item's number, if type==BLJSON_Number */

	char *string;				/* The item's name string, if this item is the child of, or is in the list of subitems of an object. */
} BLJSON;

typedef struct BLJSON_Hooks {
      void *(*malloc_fn)(size_t sz);
      void (*free_fn)(void *ptr);
} BLJSON_Hooks;

/* Supply malloc, realloc and free functions to BLJSON */
extern void BLJSON_InitHooks(BLJSON_Hooks* hooks);


/* Supply a block of JSON, and this returns a BLJSON object you can interrogate. Call BLJSON_Delete when finished. */
extern BLJSON *BLJSON_Parse(const char *value);
/* Render a BLJSON entity to text for transfer/storage. Free the char* when finished. */
extern char  *BLJSON_Print(BLJSON *item);
/* Render a BLJSON entity to text for transfer/storage without any formatting. Free the char* when finished. */
extern char  *BLJSON_PrintUnformatted(BLJSON *item);
/* Delete a BLJSON entity and all subentities. */
extern void   BLJSON_Delete(BLJSON *c);

/* Returns the number of items in an array (or object). */
extern int	  BLJSON_GetArraySize(BLJSON *array);
/* Retrieve item number "item" from array "array". Returns NULL if unsuccessful. */
extern BLJSON *BLJSON_GetArrayItem(BLJSON *array,int item);
/* Get item "string" from object. Case insensitive. */
extern BLJSON *BLJSON_GetObjectItem(BLJSON *object,const char *string);

/* For analysing failed parses. This returns a pointer to the parse error. You'll probably need to look a few chars back to make sense of it. Defined when BLJSON_Parse() returns 0. 0 when BLJSON_Parse() succeeds. */
extern const char *BLJSON_GetErrorPtr(void);
	
/* These calls create a BLJSON item of the appropriate type. */
extern BLJSON *BLJSON_CreateNull(void);
extern BLJSON *BLJSON_CreateTrue(void);
extern BLJSON *BLJSON_CreateFalse(void);
extern BLJSON *BLJSON_CreateBool(int b);
extern BLJSON *BLJSON_CreateNumber(double num);
extern BLJSON *BLJSON_CreateString(const char *string);
extern BLJSON *BLJSON_CreateArray(void);
extern BLJSON *BLJSON_CreateObject(void);

/* These utilities create an Array of count items. */
extern BLJSON *BLJSON_CreateIntArray(const int *numbers,int count);
extern BLJSON *BLJSON_CreateFloatArray(const float *numbers,int count);
extern BLJSON *BLJSON_CreateDoubleArray(const double *numbers,int count);
extern BLJSON *BLJSON_CreateStringArray(const char **strings,int count);

/* Append item to the specified array/object. */
extern void BLJSON_AddItemToArray(BLJSON *array, BLJSON *item);
extern void	BLJSON_AddItemToObject(BLJSON *object,const char *string,BLJSON *item);
/* Append reference to item to the specified array/object. Use this when you want to add an existing BLJSON to a new BLJSON, but don't want to corrupt your existing BLJSON. */
extern void BLJSON_AddItemReferenceToArray(BLJSON *array, BLJSON *item);
extern void	BLJSON_AddItemReferenceToObject(BLJSON *object,const char *string,BLJSON *item);

/* Remove/Detatch items from Arrays/Objects. */
extern BLJSON *BLJSON_DetachItemFromArray(BLJSON *array,int which);
extern void   BLJSON_DeleteItemFromArray(BLJSON *array,int which);
extern BLJSON *BLJSON_DetachItemFromObject(BLJSON *object,const char *string);
extern void   BLJSON_DeleteItemFromObject(BLJSON *object,const char *string);
	
/* Update array items. */
extern void BLJSON_ReplaceItemInArray(BLJSON *array,int which,BLJSON *newitem);
extern void BLJSON_ReplaceItemInObject(BLJSON *object,const char *string,BLJSON *newitem);

/* Duplicate a BLJSON item */
extern BLJSON *BLJSON_Duplicate(BLJSON *item,int recurse);
/* Duplicate will create a new, identical BLJSON item to the one you pass, in new memory that will
need to be released. With recurse!=0, it will duplicate any children connected to the item.
The item->next and ->prev pointers are always zero on return from Duplicate. */

/* ParseWithOpts allows you to require (and check) that the JSON is null terminated, and to retrieve the pointer to the final byte parsed. */
extern BLJSON *BLJSON_ParseWithOpts(const char *value,const char **return_parse_end,int require_null_terminated);

extern void BLJSON_Minify(char *json);

/* Macros for creating things quickly. */
#define BLJSON_AddNullToObject(object,name)		BLJSON_AddItemToObject(object, name, BLJSON_CreateNull())
#define BLJSON_AddTrueToObject(object,name)		BLJSON_AddItemToObject(object, name, BLJSON_CreateTrue())
#define BLJSON_AddFalseToObject(object,name)		BLJSON_AddItemToObject(object, name, BLJSON_CreateFalse())
#define BLJSON_AddBoolToObject(object,name,b)	BLJSON_AddItemToObject(object, name, BLJSON_CreateBool(b))
#define BLJSON_AddNumberToObject(object,name,n)	BLJSON_AddItemToObject(object, name, BLJSON_CreateNumber(n))
#define BLJSON_AddStringToObject(object,name,s)	BLJSON_AddItemToObject(object, name, BLJSON_CreateString(s))

/* When assigning an integer value, it needs to be propagated to valuedouble too. */
#define BLJSON_SetIntValue(object,val)			((object)?(object)->valueint=(object)->valuedouble=(val):(val))

#ifdef __cplusplus
}
#endif

#endif
