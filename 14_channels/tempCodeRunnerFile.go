	wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func ()  {
			ch <- 42
			wg.Done()
		}()