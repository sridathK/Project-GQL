package repository

import custommodel "project-gql/models"

type Company interface {
	CreateCompany(custommodel.Company) (custommodel.Company, error)
	GetAllCompany() ([]custommodel.Company, error)
	GetCompany(id int) (custommodel.Company, error)
	CreateJob(j custommodel.Job) (custommodel.Job, error)
	GetJobs(id int) ([]custommodel.Job, error)
	GetAllJobs() ([]custommodel.Job, error)
}

func (r *Repo) CreateCompany(u custommodel.Company) (custommodel.Company, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return custommodel.Company{}, err
	}
	return u, nil
}

func (r *Repo) GetAllCompany() ([]custommodel.Company, error) {
	var s []custommodel.Company
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Repo) GetCompany(id int) (custommodel.Company, error) {
	var m custommodel.Company

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return custommodel.Company{}, err
	}
	return m, nil

}

func (r *Repo) CreateJob(j custommodel.Job) (custommodel.Job, error) {
	err := r.db.Create(&j).Error
	if err != nil {
		return custommodel.Job{}, err
	}
	return j, nil
}

func (r *Repo) GetJobs(id int) ([]custommodel.Job, error) {
	var m []custommodel.Job

	tx := r.db.Where("uid = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil

}

func (r *Repo) GetAllJobs() ([]custommodel.Job, error) {
	var s []custommodel.Job
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}
