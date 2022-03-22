function sendRequest(method, url, body = null) {
    const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
    };

    return fetch(url, {
        method: method,
        body: JSON.stringify(body),
        headers: headers
    }).then(async (response) => ({
        status: response.status,
        data: await response.json(),
    }));
};

function onDateSelect() {
    const selectedDate = new Date(dateSelect.value);

    dateSelect.hidden = !dateSelect.hidden;

    if (isNaN(selectedDate)) return;

    printWeek(selectedDate);
}

const dateSelect = document.querySelector('.date-select');

function showDateSelect() {
    dateSelect.hidden = !dateSelect.hidden;
}

function printWeek(date = new Date()) {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();

    sendRequest('POST', 'https://gumrf.xyz/api/week.what-a-week', {
        year: year,
        month: month,
        day: day
    }).then(data => {
        options = {
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        }

        const stringDate = date.toLocaleString('ru', options);
        const week = data.data.week_fraction == 'numerator' ? 'Числитель' : 'Знаменатель';

        let whatWeekText = '';
        if (date.setHours(0, 0, 0, 0) == new Date().setHours(0, 0, 0, 0)) whatWeekText = 'Сегодня';
        whatWeekText += `\n${stringDate}\n${week}`;

        document.querySelector('.what-week-text').innerText = whatWeekText;
    });
}

printWeek();

dateSelect.addEventListener('change', onDateSelect);
document.querySelector('.custom-week').onclick = showDateSelect;
