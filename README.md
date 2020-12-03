# REST service for "Subscription as a Service"

## API Service

The primary aspect of this programming challenge is to implement the following two APIs:

- `/user`
- `/subscription`

The details of each of these APIs are as follows:

### `/user`

This is a simple CRUD API that adds a user to DB.

#### PUT `/user/<username>`

- creates a user with specified username in the DB.

**Sample Input** : PUT `/user/jay`

**Required Output** :

- Just a HTTP status: 200 on success, other appropriate code for failures

#### GET `/user/<username>`

**Sample Input** : GET `/user/jay`

**Sample Output** :

```
{
	"user_name": "jay",
	"created_at": "2020-02-29 19:30:00"
}
```

- Please make sure that the output is in the format above to allow our automated test cases to pass.

### `/subscription`

- This is the primary API being tested in this challenge.
- This will need to provide mechanisms to:
  - Register a new subscription for an existing user, with a specified plan and start date

#### POST `/subscription/`

**Inputs** :

```
{
	"user_name": <username string>,
	"plan_id": <plan_id string>,
	"start_date": <date string in YYYY-MM-DD format>
}
```

**Sample Input**

```
{
	"user_name": "jay",
	"plan_id": "PRO_1M",
	"start_date": "2020-03-03"
}
```

**Expected Output** :

```
{
	"status": <["SUCCESS"|"FAILIURE"]>,
	"amount": <+/- amount credited/debited>
}
```

**Sample Output** :

```
{
	"status": "SUCCESS",
	"amount": -200
}
```

**Additional details** :

- API output shall conform to the above described format so that it passes our automated testing.
- On success, return 200 HTTP status. For failures, pick an appropriate HTTP code.
- The timestamp indicates the start date for the new plan, and it will be valid for the number
  of days shown in the table below.
- plan_id can be one of those listed in the table below:

| Plan ID | Validity (in days) | Cost (USD) |
| ------- | ------------------ | ---------- |
| FREE    | Infinite           | 0          |
| TRIAL   | 7                  | 0          |
| LITE_1M | 30                 | 100        |
| PRO_1M  | 30                 | 200        |
| LITE_6M | 180                | 500        |
| PRO_6M  | 180                | 900        |

- The service is expected to check if the new plan addition entails an upgrade of plans or a downgrade.
  - If it is an upgrade, the service must make a call to the Payment API server (see below), with a debit of the amount applicable for the upgrade.
  - If it is a downgrade, make a credit request with the appropriate amount to the Payment API server.
- Once payment succeeds, this service shall make the necessary changes in database to update the user's subscription plan.
- The `amount` field in the API response shall be negative to indicate debits and positive in case
  of credits.

#### GET `/subscription/<username>/<date>`

**Sample Input** : `/subscription/jay/2020-02-29`

Note that the date in the above is optional:
`/subscription/jay`

**Expected Output** :

- When input date is specified
  - plan_id that will be active for user at specified date.
  - Number of days left in plan from the specified input date
  - **Sample output**

```
{
	"plan_id": "PRO_1M",
	"days_left": 3
}
```

- When input date is NOT specified
  - List all subscription entries available in database for user with start and valid till dates.
  - **Sample output**

```
[
	{
		"plan_id": "TRIAL",
		"start_date": "2020-02-22",
		"valid_till": "2020-02-28"
	},
	{
		"plan_id": "PRO_1M",
		"start_date": "2020-02-29",
		"valid_till": "2020-03-30"
	}
]
```

**Additional points of Note** :

- API output shall conform to the above described format so that it passes our automated
  testing.

## Payment API

This external service provides a single API endpoint described below:

**End point** : `/payment`

**Request type** : POST

**Request body** :

```
{
	"user_name": <string>,
	"payment_type": <one of "DEBIT"|"CREDIT">,
	"amount": <number>
}
```

**Response body** :

```
{
	"payment_id": <uuid> // eg. "24242-3443-sdstg-3343",
	"status": <one of "SUCCESS"|"FAILIURE">
}
```

**Additional points of Note** :

- This service is implemented so that it intentionally errors out sometimes (approx 25% calls
  fail). This failure needs to be handled appropriately in the subscriptions service.
