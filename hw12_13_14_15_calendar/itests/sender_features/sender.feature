# file: features/sender.feature

# http://localhost:8081/
# http://calendar_service:8081/

Feature: notification sender

    Scenario: publish notification
		When I publish notification with data:
		"""
		{
			"ID": "1",
			"UserID": 1,
			"Title": "test@test.ru"
		}
		"""
		Then The notification should be "test@test.ru"