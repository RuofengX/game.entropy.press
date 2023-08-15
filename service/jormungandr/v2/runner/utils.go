package runner

import (
	"errors"
	"math"
	"jormungandr/v2/proto/base"
)

// 获取两个实体间的距离
func getEntityDistance(e1 *base.Entity, e2 *base.Entity) float64{
	return math.Sqrt(math.Pow(e1.Velo.X-e2.Velo.X, 2) + math.Pow(e1.Velo.Y-e2.Velo.Y, 2))
}

// 获取给出的两个ID对应的实体之间的距离
// 获取s空间中，编号为one和other的两个实体之间的距离
// 两个实体都需要有.Velo.X和.Velo.Y属性，否则抛出异常
// 返回：单精度的距离，错误
func GetDistance(s *base.Space, one uint64, other uint64) (float64, error) {
	e1, ok := s.Entity[one]
	if !ok {
		return 0.0, errors.New("第一位实体不存在")
	}

	e2, ok := s.Entity[other]
	if !ok {
		return 0.0, errors.New("第二位实体不存在")
	}
	rtn := getEntityDistance(e1, e2)
	return rtn, nil
}

// 获取附近的实体
// 获取s空间中，编号为one的实体周围radius范围的实体列表
// 返回：实体列表、错误
// TODO: 需要对其性能测试
// TODO: 最好加入缓存机制
func GetNearby(s *base.Space, one uint64, radius float64) ([]uint64, error) {
	e1, ok := s.Entity[one]
	if !ok {
		return []uint64{}, errors.New("实体不存在")
	}

	rtn := []uint64{}

	for index, e2 := range s.Entity {
		dis := getEntityDistance(e1, e2)
		if dis <= radius {
			rtn = append(rtn, index)
		}
	}
	return rtn, nil
}
