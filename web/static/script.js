// Слушатель события загрузки DOM.
document.addEventListener('DOMContentLoaded', async function() {
    // Попытка инициализации particles.js.
    try {
        // Загрузка конфигурации particles.js из файла particles.json.
        // Убедитесь, что файл particles.json находится в доступном месте относительно HTML-файла.
        particlesJS.load('particles-js', 'particles.json', function() {
            console.log('callback - Конфигурация particles.js успешно загружена.');
        });
    } catch (error) {
        // Логирование ошибки, если конфигурация не может быть загружена.
        console.error('Ошибка при загрузке конфигурации particles.js:', error);
    }

    // Получение элемента формы и кнопки отправки.
    const form = document.getElementById('uidForm');
    const submitButton = form.querySelector('button[type="submit"]');
    // Сохранение оригинального текста кнопки для последующего восстановления.
    const originalButtonText = submitButton.textContent;

    // Обработчик события отправки формы.
    form.addEventListener('submit', async function(e) {
        e.preventDefault(); // Предотвращение стандартного поведения формы.

        // Изменение текста кнопки на время отправки и её отключение.
        submitButton.textContent = 'Отправка...';
        submitButton.disabled = true;

        // Создание объекта FormData из формы для отправки.
        const formData = new FormData(form);

        try {
            // Отправка данных формы на сервер методом POST.
            const response = await fetch(form.action, {
                method: 'POST',
                body: formData
            });

            // Проверка статуса ответа от сервера.
            if (response.ok) {
                console.log('Данные формы успешно отправлены.');
                alert('Данные успешно отправлены!');
            } else {
                console.error('Ошибка при отправке данных формы. Статус ответа:', response.status);
                alert('Не удалось отправить данные.');
            }
        } catch (error) {
            // Логирование ошибки при отправке данных.
            console.error('Ошибка при отправке данных формы:', error);
            alert('Произошла ошибка при отправке. Пожалуйста, попробуйте снова.');
        } finally {
            // Восстановление оригинального текста кнопки и её активации после отправки.
            submitButton.textContent = originalButtonText;
            submitButton.disabled = false;
        }
    });
});