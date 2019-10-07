package crud

import (
	"blogos/models"

	"github.com/jinzhu/gorm"
)

type repositoryPostCRUD struct {
	db *gorm.DB
}

func NewRepositoryPostCRUD(db *gorm.DB) *repositoryPostCRUD {
	return &repositoryPostCRUD{db}
}

func (r *repositoryPostCRUD) Save(post models.Post) (models.Post, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(done)
		if err = r.db.Debug().Model(&models.Post{}).Create(&post).Error; err != nil {
			done <- false
			return
		}
		done <- true
	}(done)
	if res := <-done; res {
		return post, nil
	} else {
		return models.Post{}, err
	}
}

func (r *repositoryPostCRUD) FindAll() ([]models.Post, error) {
	var err error
	done := make(chan bool)
	var posts []models.Post
	go func(ch chan<- bool) {
		defer close(done)
		if err = r.db.Debug().Model(&models.Post{}).Limit(10).Find(&posts).Error; err != nil {
			done <- false
			return
		}
		jobs := make(chan int)
		results := make(chan int)
		for j := 0; j < 2; j++ {
			go func(jobs <-chan int, results chan<- int, done chan bool) {
				for i := range jobs {
					if err := r.db.Model(&posts[i]).Related(&posts[i].Author, "author_id").Error; err != nil {
						done <- false
						return
					}
					results <- 1
				}
			}(jobs, results, done)
		}
		go func() {
			for i, _ := range posts {
				jobs <- i
			}
			close(jobs)
		}()
		for _, _ = range posts {
			<-results
		}
		done <- true
	}(done)
	if res := <-done; res {
		return posts, nil
	} else {
		return []models.Post{}, err
	}
}

func (r *repositoryPostCRUD) FindById(pid uint) (models.Post, error) {
	var err error
	done := make(chan bool)
	var post models.Post
	go func(ch chan<- bool) {
		defer close(done)
		if err = r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).Take(&post).Error; err != nil {
			done <- false
			return
		}
		if err := r.db.Model(&post).Related(&post.Author, "author_id").Error; err != nil {
			done <- false
			return
		}
		done <- true
	}(done)
	if res := <-done; res {
		return post, nil
	} else {
		return models.Post{}, err
	}
}
