package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

// Aлгоритм реализует стратегию Round-Robin балансировки нагрузки, обеспечивая
// передачу запросов по кругу, проверяя контекст на предмет отмены и автоматически
// пробуя альтернативные бекенды, если первый выбранный оказался недоступен.
// Код также аккуратно учитывает ситуацию, когда все бекенды оказываются недоступны,
// формируя соответствующее сообщение об ошибке.

func (b *Balancer) Invoke(ctx context.Context, req Request) (Response, error) {
	// Получение количества бекендов
	n := len(b.backends)

	// Алгоритм выбирает следующий бекенд, основываясь на предыдущем индексе.
	// Переменная next увеличивается атомарно, чтобы предотвратить гонку данных
	// при доступе из нескольких горутин. Затем производится деление по модулю,
	// чтобы обеспечить круговой перебор индексов бекендов.
	start := int(atomic.AddUint64(&b.next, 1)-1) % n

	var lastErr error

	// Перебираем все бекенды
	for i := 0; i < n; i++ {

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// Обращаемся к бекенду
		idx := (start + i) % n
		resp, err := b.backends[idx].Invoke(ctx, req)
		if err == nil {
			return resp, nil
		}

		lastErr = err
	}
	return nil, fmt.Errorf("all %d backends failed: %w", n, lastErr)
}
