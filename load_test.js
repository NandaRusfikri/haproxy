import http from 'k6/http'
import { check, sleep } from 'k6'
export const options = {
    vus: 1000,
    duration: '1s',
  };


export default function () {
  let res = http.get('http://localhost')

  check(res, { 'success request': (r) => r.status === 200 })

  sleep(0.3)
}
