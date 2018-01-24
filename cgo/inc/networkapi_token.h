#ifndef NetworkAPI__networkapi_token__H__
#define NetworkAPI__networkapi_token__H__

#include <sys/time.h>
#include <pthread.h>

#ifdef __cplusplus
extern "C"
{
#endif

#define TOKEN_GRANULARITY	(1000)

/*
 * simple token-bucket implementation,
 * When query the token, current tokens is increased time interval * tb->quota / tb->interval
 * cost is  TOKEN_BUCKET_GRANULARITY;
 */
struct networkapi_token_bucket_t {
    /*mutex*/
    pthread_mutex_t mutex;
	/*last querry jiffie*/
	struct timeval	last_query;
	/*current tokens*/
	int		tokens;
	/*max tokens allowed*/
	int		burst;
	int		step;
};

#ifndef networkapi_timersub
#define	networkapi_timersub(tvp, uvp, vvp)						\
	do {								\
		(vvp)->tv_sec = (tvp)->tv_sec - (uvp)->tv_sec;		\
		(vvp)->tv_usec = (tvp)->tv_usec - (uvp)->tv_usec;	\
		if ((vvp)->tv_usec < 0) {				\
			(vvp)->tv_sec--;				\
			(vvp)->tv_usec += 1000000;			\
		}							\
	} while (0)
#endif

extern int networkapi_token_init(struct networkapi_token_bucket_t *tb, int quota, int interval);
extern int networkapi_token_bucket_query(struct networkapi_token_bucket_t *tb);
        
#ifdef __cplusplus
}
#endif

#endif
