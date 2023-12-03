/**
 * Generated by orval v6.21.0 🍺
 * Do not edit manually.
 * stlatica_internal_api
 * stlatica internal api
 * OpenAPI spec version: 0.1.0
 */
import {
  faker
} from '@faker-js/faker'
import {
  HttpResponse,
  delay,
  http
} from 'msw'

export const getGetPlatMock = () => ({content: faker.word.sample(), plat_id: faker.word.sample()})

export const getPlatMock = () => [
http.get('*/internal/v1/plats/:platId', async () => {
        await delay(1000);
        return new HttpResponse(JSON.stringify(getGetPlatMock()),
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),http.delete('*/internal/v1/plats/:platId', async () => {
        await delay(1000);
        return new HttpResponse(null,
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),http.post('*/internal/v1/plats', async () => {
        await delay(1000);
        return new HttpResponse(null,
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),]
