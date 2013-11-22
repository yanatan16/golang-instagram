package instagram

// IterateMedia makes pagination easy by converting the repeated api.NextMedias() call to a channel of media.
// Media will be passed in the reverse order of individual requests, for instance GetUserRecentMedia will go in reverse CreatedTime order.
// If you desire to break early, pass in a doneChan and close when you are breaking from iteration
func (api *Api) IterateMedia(res *PaginatedMediasResponse, doneChan <-chan bool) (<-chan *Media, <-chan error) {
	mediaChan := make(chan *Media)
	errChan := make(chan error, 1)

	go func() {
		defer close(mediaChan)
		defer close(errChan)

		for {
			if res == nil {
				return
			}

			if len(res.Medias) == 0 {
				// No more Medias
				return
			}

			// Iterate backwards
			for i := len(res.Medias) - 1; i >= 0; i-- {
				select {
				case mediaChan <- &res.Medias[i]:
				case <-doneChan:
					return
				}
			}

			// Paginate to next response
			var err error
			res, err = api.NextMedias(res.Pagination)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	return mediaChan, errChan
}

// IterateUsers makes pagination easy by converting the repeated api.NextUsers() call to a channel of users.
// Users will be passed in the reverse order of individual requests.
// If you desire to break early, pass in a doneChan and close when you are breaking from iteration
func (api *Api) IterateUsers(res *PaginatedUsersResponse, doneChan <-chan bool) (<-chan *User, <-chan error) {
	userChan := make(chan *User)
	errChan := make(chan error)

	go func() {
		defer close(userChan)
		defer close(errChan)

		for {
			if res == nil {
				return
			}

			if len(res.Users) == 0 {
				// No more users
				return
			}

			// Iterate backwards
			for i := len(res.Users) - 1; i >= 0; i-- {
				select {
				case userChan <- &res.Users[i]:
				case <-doneChan:
					return
				}
			}

			// Paginate to next response
			var err error
			res, err = api.NextUsers(res.Pagination)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	return userChan, errChan
}
