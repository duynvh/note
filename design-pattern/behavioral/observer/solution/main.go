package main

import "fmt"

type Observer interface {
	ReceiveNotify(j Job)
}

type Developer struct{}

func (Developer) ReceiveNotify(j Job) {
	fmt.Println("Many thanks, I've received job:", j.Title)
}

type Job struct {
	Title string
}

// Aka: Subject
type ITJobsCompany struct {
	jobs      []Job
	observers []Observer
}

func (com *ITJobsCompany) AddObserver(o Observer) {
	com.observers = append(com.observers, o)
}

func (com *ITJobsCompany) RemoveObserver(o Observer) {
	for i := range com.observers {
		if com.observers[i] == o {
			com.observers = append(com.observers[:i], com.observers[i+1:]...)
			return
		}
	}
}

func (com *ITJobsCompany) notifyToObservers(j Job) {
	for i := range com.observers {
		com.observers[i].ReceiveNotify(j)
	}
}

func (com *ITJobsCompany) AddJob(j Job) {
	com.jobs = append(com.jobs, j)
	com.notifyToObservers(j)
}

func main() {
	itComp := ITJobsCompany{}
	developer := Developer{}

	// Developer will be added an observer
	itComp.AddObserver(developer)

	// Developer will receive new jobs
	itComp.AddJob(Job{Title: "Go"})
	itComp.AddJob(Job{Title: "Java"})

	// Developer no longer receive new job
	itComp.RemoveObserver(developer)

	itComp.AddJob(Job{Title: "Python"})
}
