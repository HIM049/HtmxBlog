package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"sort"
	"strconv"
	"strings"
)

type StatsSummary struct {
	TotalPV      int64
	TotalUV      int64
	PopularPosts []PostStat
	TopReferers  []RefererStat
}

type PostStat struct {
	Rank   int
	PostID uint
	Title  string
	Views  int64
}

type RefererStat struct {
	Rank    int
	Referer string
	Count   int64
}

type PathCount struct {
	Path  string
	Count int64
}

type RefererCount struct {
	Referer string
	Count   int64
}

func GetStats() (*StatsSummary, error) {
	// 1. Total PV
	var totalPV int64
	if err := config.DB.Model(&model.AccessRecord{}).Count(&totalPV).Error; err != nil {
		return nil, err
	}

	// 2. Total UV
	var totalUV int64
	if err := config.DB.Model(&model.AccessRecord{}).Distinct("ip").Count(&totalUV).Error; err != nil {
		return nil, err
	}

	// 3. Popular Posts
	var pathCounts []PathCount
	err := config.DB.Model(&model.AccessRecord{}).
		Select("path, count(*) as count").
		Where("path LIKE ?", "/post/%").
		Group("path").
		Order("count DESC").
		Find(&pathCounts).Error
	if err != nil {
		return nil, err
	}

	var popularPosts []PostStat
	var postIDs []uint
	postViews := make(map[uint]int64)

	for _, pc := range pathCounts {
		parts := strings.Split(strings.Trim(pc.Path, "/"), "/")
		if len(parts) == 2 && parts[0] == "post" {
			id, err := strconv.Atoi(parts[1])
			if err == nil {
				postViews[uint(id)] += pc.Count
				postIDs = append(postIDs, uint(id))
			}
		}
	}

	if len(postIDs) > 0 {
		var posts []model.Post
		if err := config.DB.Where("id IN ?", postIDs).Find(&posts).Error; err == nil {
			postMap := make(map[uint]string)
			for _, p := range posts {
				postMap[p.ID] = p.Title
			}
			for id, count := range postViews {
				title, ok := postMap[id]
				if !ok {
					title = "Deleted Post (ID " + strconv.Itoa(int(id)) + ")"
				}
				popularPosts = append(popularPosts, PostStat{
					PostID: id,
					Title:  title,
					Views:  count,
				})
			}
		}
	}

	sort.Slice(popularPosts, func(i, j int) bool {
		return popularPosts[i].Views > popularPosts[j].Views
	})
	for i := range popularPosts {
		popularPosts[i].Rank = i + 1
	}

	// 4. Referers
	var topReferers []RefererStat
	err = config.DB.Model(&model.AccessRecord{}).
		Select("referer, count(*) as count").
		Group("referer").
		Order("count DESC").
		Scan(&topReferers).Error
	if err != nil {
		return nil, err
	}

	for i := range topReferers {
		topReferers[i].Rank = i + 1
		if strings.TrimSpace(topReferers[i].Referer) == "" {
			topReferers[i].Referer = "Other"
		}
	}

	return &StatsSummary{
		TotalPV:      totalPV,
		TotalUV:      totalUV,
		PopularPosts: popularPosts,
		TopReferers:  topReferers,
	}, nil
}
