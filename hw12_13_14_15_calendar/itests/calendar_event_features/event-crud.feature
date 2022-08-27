# file: features/event-crud.feature

# http://localhost:8081/
# http://calendar_service:8081/

Feature: API calendar event crud

    Scenario: Calendar event create
		When I send "POST" request to "http://calendar_service:8081/create-event" with "application/json" data:
		"""
		{
			"user_id": 2,
			"title": "test@test.ru",
			"duration": "3 days 04:05:06"
		}
		"""
		Then The response code should be 200

    Scenario: Calendar event get
		When I send "GET" request to "http://calendar_service:8081/get-user-events/2"
		Then The response code should be 200
		And the response should match json:
		"""
        {
            "Content":[{"id":"2","userId":"2","title":"test@test.ru"}]
        }
        """

    Scenario: Calendar event delete
		When I send "POST" request to "http://calendar_service:8081/delete-event" with "application/json" data:
        """
		{
			"id": "2"
		}
		"""
		Then The response code should be 200