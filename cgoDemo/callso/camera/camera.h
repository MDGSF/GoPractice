#ifndef __CAMERA_FORGO_H__
#define __CAMERA_FORGO_H__
#ifdef __cplusplus
extern "C" {
#endif

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdint.h>


/*
 * @brief: startCamera run camera module, it will get framge from camera continuous.
 * @default value: 1, 1000, 1, 1000, 704, 576, 25, 10
 * @param adas[in]: 1
 * @param adasCacheSize[in]: 1000
 * @param dsm[in]: 1
 * @param dsmCacheSize[in]: 1000
 * @param width[in]: 704
 * @param height[in]: 576
 * @param quality[in]: 25
 * @param fps[in]: 10
 **/
void startCamera(
    int adas,
	int adasCacheSize,
    int dsm,
	int dsmCacheSize,
    int width,
    int height,
    int quality,
    int fps
);

/*
 * @brief: get latest adas image.
 * @param piImageSize[out]: the image size return.
 * @param ppcImageData[out]: the image binary data return.
 * @return: EXIT_SUCCESS, else failed.
 */
int getLatestAdasImage(int * piImageSize, char ** ppcImageData);

/*
 * @brief: get latest dsm image.
 * @param piImageSize[out]: the image size return.
 * @param ppcImageData[out]: the image binary data return.
 * @return: EXIT_SUCCESS, else failed.
 */
int getLatestDsmImage(int * piImageSize, char ** ppcImageData);

/*
 * @brief: get adas video from startTime to endTime.
 * @param startTime[in]: the video start time, in second.
 * @param endTime[in]: the video end time, in second.
 * @param piVideoSize[out]: the video size return.
 * @param ppcVideoData[out]: the video binary data return.
 * @return: EXIT_SUCCESS, else failed.
 */
int getAdasVideo(int64_t startTime, int64_t endTime, int * piVideoSize, char ** ppcVideoData);

/*
 * @brief: get dsm video from startTime to endTime.
 * @param startTime[in]: the video start time, in second.
 * @param endTime[in]: the video end time, in second.
 * @param piVideoSize[out]: the video size return.
 * @param ppcVideoData[out]: the video binary data return.
 * @return: EXIT_SUCCESS, else failed.
 */
int getDsmVideo(int64_t startTime, int64_t endTime, int * piVideoSize, char ** ppcVideoData);

#ifdef __cplusplus
}
#endif
#endif
