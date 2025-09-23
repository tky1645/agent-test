# Requirements Document

## Introduction

フロントエンドアプリケーションを実際のバックエンドAPIと連携させるための機能を実装します。現在はモックデータとlocalStorageを使用していますが、これを実際のHTTP APIコールに置き換え、認証、エラーハンドリング、ローディング状態管理を含む包括的なAPI統合を提供します。

## Requirements

### Requirement 1

**User Story:** As a frontend developer, I want to replace mock API services with real HTTP API calls, so that the application can communicate with the actual backend server.

#### Acceptance Criteria

1. WHEN the application starts THEN the API service SHALL use HTTP requests instead of localStorage mock data
2. WHEN making API calls THEN the system SHALL include proper authentication headers with JWT tokens
3. WHEN API calls fail THEN the system SHALL handle errors gracefully and provide meaningful error messages
4. WHEN API calls are in progress THEN the system SHALL show appropriate loading states to users

### Requirement 2

**User Story:** As a user, I want the application to authenticate with the backend API, so that I can securely access my plant data.

#### Acceptance Criteria

1. WHEN a user logs in THEN the system SHALL send credentials to the backend authentication endpoint
2. WHEN authentication succeeds THEN the system SHALL store the JWT token securely
3. WHEN making authenticated requests THEN the system SHALL include the Authorization header with Bearer token
4. WHEN the token expires THEN the system SHALL redirect users to login and clear stored credentials

### Requirement 3

**User Story:** As a user, I want to manage my plants through the backend API, so that my data is persisted on the server.

#### Acceptance Criteria

1. WHEN fetching plants THEN the system SHALL call GET /api/plants with authentication
2. WHEN creating a plant THEN the system SHALL call POST /api/plants with plant data and authentication
3. WHEN updating plant information THEN the system SHALL call PUT /api/plants/{id} with updated data
4. WHEN deleting a plant THEN the system SHALL call DELETE /api/plants/{id} with authentication

### Requirement 4

**User Story:** As a user, I want to record watering activities through the backend API, so that my watering history is accurately tracked.

#### Acceptance Criteria

1. WHEN recording a watering event THEN the system SHALL call POST /api/plants/{id}/watering with timestamp and notes
2. WHEN fetching watering history THEN the system SHALL call GET /api/plants/{id}/watering-records
3. WHEN getting plant status THEN the system SHALL call GET /api/plants/{id}/status for current watering status
4. WHEN watering records are updated THEN the UI SHALL reflect the latest data immediately

### Requirement 5

**User Story:** As a user, I want proper error handling and feedback, so that I understand what happens when API calls fail.

#### Acceptance Criteria

1. WHEN network errors occur THEN the system SHALL display user-friendly error messages
2. WHEN server returns validation errors THEN the system SHALL show specific field-level error messages
3. WHEN authentication fails THEN the system SHALL redirect to login with appropriate error message
4. WHEN API calls timeout THEN the system SHALL provide retry options to users

### Requirement 6

**User Story:** As a user, I want responsive loading states, so that I know when the application is processing my requests.

#### Acceptance Criteria

1. WHEN API calls are in progress THEN the system SHALL show loading spinners or skeleton screens
2. WHEN submitting forms THEN buttons SHALL be disabled and show loading state
3. WHEN fetching data THEN the system SHALL show appropriate loading indicators
4. WHEN operations complete THEN loading states SHALL be cleared immediately

### Requirement 7

**User Story:** As a developer, I want configurable API endpoints, so that the application can work with different backend environments.

#### Acceptance Criteria

1. WHEN deploying to different environments THEN the API base URL SHALL be configurable via environment variables
2. WHEN API endpoints change THEN the configuration SHALL be centralized and easy to update
3. WHEN debugging API calls THEN the system SHALL provide proper logging and request/response visibility
4. WHEN API versions change THEN the system SHALL support version headers or URL versioning