package core

import (
	"fmt"
	"time"

	"git.qdreads.com/gotools/redis"
)

type BaseCache struct {
	PrefixKey      string        `desc:"缓存key前缀"`
	ExpirationTime time.Duration `desc:"缓存有效期,单位:秒"`
}

func (cache BaseCache) Delete(ctx *Context, key string, instance ...string) error {
	if key == "" {
		key = cache.PrefixKey
	} else {
		key = fmt.Sprintf("%s::{%s}", cache.PrefixKey, key)
	}

	return Redis(instance...).Del(ctx, key).Err()
}

// 如果需要随机过期时间，可以重写该方法
func (cache BaseCache) GetExpirationTime() time.Duration {
	return cache.ExpirationTime
}

type StringCache struct {
	BaseCache
}

type StringFunc func() (string, error)

func (cache StringCache) Get(ctx *Context, key string, f StringFunc, instance ...string) (str string, err error) {
	if key == "" {
		key = cache.PrefixKey
	} else {
		key = fmt.Sprintf("%s::{%s}", cache.PrefixKey, key)
	}

	str, err = Redis(instance...).Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// 没有获取对应的数据
			// TODO 添加分布式锁
			if result, resStr := f(); resStr != nil {
				err = resStr
				return
			} else {
				err = Redis(instance...).Set(ctx, key, result, cache.GetExpirationTime()).Err()
				str = result
			}
		}
	}

	return
}

// hash

type HashCache struct {
	BaseCache
}

type HashFunc func(fields ...string) (data map[string]string, err error)

func (cache HashCache) Get(ctx *Context, key string, f HashFunc, fields []string, instance ...string) (data map[string]string, err error) {
	if key == "" {
		key = cache.PrefixKey
	} else {
		key = fmt.Sprintf("%s::{%s}", cache.PrefixKey, key)
	}

	data = map[string]string{}

	if len(fields) == 0 {
		if data, err = Redis(instance...).HGetAll(ctx, key).Result(); err != nil {
			return
		} else {
			if err == redis.Nil {
				if data, err = f(); err != nil {
					return
				}

				values := []interface{}{}
				for key, v := range data {
					values = append(values, key, v)
				}

				err = Redis(instance...).HSet(ctx, key, values...).Err()
				Redis(instance...).Expire(ctx, key, cache.GetExpirationTime())
				return
			}
		}
	}

	noExistFields := []string{}
	if dataList, dataErr := Redis(instance...).HMGet(ctx, key, fields...).Result(); err != nil {
		err = dataErr
		return
	} else {
		for index, value := range dataList {
			if value != nil {
				data[fields[index]] = value.(string)
			} else {
				noExistFields = append(noExistFields, fields[index])
			}
		}
	}

	if res, resErr := f(noExistFields...); resErr != nil {
		err = resErr
		return
	} else {
		if len(res) == 0 {
			return
		}

		values := []interface{}{}
		for key, v := range res {
			data[key] = v
			values = append(values, key, v)
		}

		Redis(instance...).HSet(ctx, key, values...).Err()
	}

	return
}
