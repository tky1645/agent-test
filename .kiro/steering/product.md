---
inclusion: always
---

# Product Domain & Business Rules

## Domain
Plant watering management system for home gardeners to track care schedules and watering history.

## Core Entities & Rules

### User
- Unique username and email required
- Authentication required for all operations
- Users can only access their own plants

### Plant
- Name must be unique per user
- Requires watering frequency (days between waterings)
- Belongs to single user (ownership enforced)
- Status: active (can be watered) or archived

### WateringRecord
- Immutable audit trail (cannot be deleted)
- Tracks timestamp, user, plant association
- Multiple waterings per day allowed
- "Days since watering" calculated from most recent record

## Business Logic

### Watering Status
- **Overdue**: days since watering > frequency requirement
- **Due Soon**: approaching frequency requirement
- **Recently Watered**: within frequency requirement

### Data Integrity
- Plant names unique per user
- Only plant owners can water plants
- Watering records preserve complete history
- All operations require authentication

## UX Principles
- Mobile-friendly, simple interface
- Visual status indicators for watering needs
- Quick watering action from plant list
- Chronological watering history display

## Feature Priority
1. Core: User auth, plant CRUD, watering tracking
2. UX: Responsive design, clear status indicators
3. Data: Audit trails, validation, error handling