### Задача Parallel Requests 

Функция, которая занимается бизнес логикой делает возврат 
    []*http.Respons

Однако при параллельной обработке запросов мы через конструкцию defer закрываем тело response
    defer res.Body.Close()

Имеем ситуацию, что когда слайс из указателей придет потребителю, он не сможет прочитать тела 
запросов, потому что мы их уже закрыли. Пример.

        Горутина:
        1. res, err := client.Do()     ← получаем Response
        2. defer res.Body.Close()      ← регистрируем закрытие при выходе
        3. newResCh <- res             ← отправляем в канал (БЛОКИРУЕТСЯ, если канал полон)
        4. горутина завершается        ← срабатывает defer, Body закрывается
        
        Основной поток:
        5. res := <-newResCh           ← получаем Response
        6. result = append(result, res)
        7. после сбора всех результатов — возвращаем слайс

        Пользователь функции:
        8. responses := callRequests(...)
        9. data, _ := io.ReadAll(responses[0].Body) ← ЧИТАЕМ ИЗ ЗАКРЫТОГО BODY! ❌


Решение! Лучше всего создать дополнительную структуру и в ней сохранить данные

        type Result struct {
            URL  string
            Data []byte
            Err  error
        }

        go func(url string) {
            res, err := http.DefaultClient.Do(newReq)
            if err != nil {
                resultCh <- Result{URL: url, Err: err}
                return
            }
            defer res.Body.Close() // ✅ безопасно
            
            data, err := io.ReadAll(res.Body)
            resultCh <- Result{
                URL:  url,
                Data: data,
                Err:  err,
            }
        }(url)
