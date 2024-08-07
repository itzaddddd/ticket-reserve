import http from "k6/http"
import { check, sleep } from "k6"

// max vus is 99, if more than it will occur error.
export const options = {
  vus: 100,
  iterations: 100
}

export default function () {
  let payload = JSON.stringify({
    user_id: 1,
    event_id: 1,
    number_to_reserve: 1
  })

  let url = "http://localhost:8080/api/v1/event/reserve"
  let params = {
    headers: { 'Content-type': 'application/json' }
  }

  const res = http.post(url, payload, params)

  check(res, { 
    'status was 200': r => {
      return  r.status === 200 
    }
  })
  
}