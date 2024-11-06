import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 10 },
    { duration: '20s', target: 40 },
    { duration: '5s', target: 20 },
  ],
};

export default function() {
  const url = 'http://localhost:8080/api/v1';
  let res = http.get(`${url}/books`);
  check(res, {'is status 200': (r) => r.status === 200});
  const payload = {title: 'string'};
  res = http.post('http://localhost:8080/api/v1/books', JSON.stringify(payload), {
    headers: {'Content-Type': 'application/json'},
  });
  check(res, {'is status 201': (r) => r.status === 201});
  sleep(1);
}
