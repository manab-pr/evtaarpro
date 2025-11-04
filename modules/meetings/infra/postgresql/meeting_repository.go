package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/meetings/domain/entities"
)

// MeetingRepository implements repository.MeetingRepository
type MeetingRepository struct {
	db *sql.DB
}

// NewMeetingRepository creates a new MeetingRepository
func NewMeetingRepository(db *sql.DB) *MeetingRepository {
	return &MeetingRepository{db: db}
}

// Create creates a new meeting
func (r *MeetingRepository) Create(ctx context.Context, meeting *entities.Meeting) error {
	query := `
		INSERT INTO meetings (id, room_id, title, description, organizer_id, start_time, status, max_participants, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		meeting.ID,
		meeting.RoomID,
		meeting.Title,
		meeting.Description,
		meeting.OrganizerID,
		meeting.StartTime,
		meeting.Status,
		meeting.MaxParticipants,
		meeting.CreatedAt,
		meeting.UpdatedAt,
	)

	return err
}

// GetByID retrieves a meeting by ID
func (r *MeetingRepository) GetByID(ctx context.Context, id string) (*entities.Meeting, error) {
	query := `
		SELECT id, room_id, title, description, organizer_id, start_time, end_time, status, jitsi_room_url, recording_url, max_participants, created_at, updated_at
		FROM meetings
		WHERE id = $1
	`

	meeting := &entities.Meeting{}
	var endTime sql.NullTime
	var jitsiURL, recordingURL sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&meeting.ID,
		&meeting.RoomID,
		&meeting.Title,
		&meeting.Description,
		&meeting.OrganizerID,
		&meeting.StartTime,
		&endTime,
		&meeting.Status,
		&jitsiURL,
		&recordingURL,
		&meeting.MaxParticipants,
		&meeting.CreatedAt,
		&meeting.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("meeting not found")
		}
		return nil, err
	}

	if endTime.Valid {
		meeting.EndTime = &endTime.Time
	}
	if jitsiURL.Valid {
		meeting.JitsiRoomURL = jitsiURL.String
	}
	if recordingURL.Valid {
		meeting.RecordingURL = recordingURL.String
	}

	return meeting, nil
}

// List retrieves meetings with pagination
func (r *MeetingRepository) List(ctx context.Context, page, pageSize int, userID string) ([]*entities.Meeting, int64, error) {
	offset := (page - 1) * pageSize

	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM meetings`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Get meetings
	query := `
		SELECT id, room_id, title, description, organizer_id, start_time, end_time, status, jitsi_room_url, recording_url, max_participants, created_at, updated_at
		FROM meetings
		ORDER BY start_time DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	meetings := make([]*entities.Meeting, 0)
	for rows.Next() {
		meeting := &entities.Meeting{}
		var endTime sql.NullTime
		var jitsiURL, recordingURL sql.NullString

		if err := rows.Scan(
			&meeting.ID,
			&meeting.RoomID,
			&meeting.Title,
			&meeting.Description,
			&meeting.OrganizerID,
			&meeting.StartTime,
			&endTime,
			&meeting.Status,
			&jitsiURL,
			&recordingURL,
			&meeting.MaxParticipants,
			&meeting.CreatedAt,
			&meeting.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}

		if endTime.Valid {
			meeting.EndTime = &endTime.Time
		}
		if jitsiURL.Valid {
			meeting.JitsiRoomURL = jitsiURL.String
		}
		if recordingURL.Valid {
			meeting.RecordingURL = recordingURL.String
		}

		meetings = append(meetings, meeting)
	}

	return meetings, total, nil
}

// ListByOrganizer retrieves meetings by organizer
func (r *MeetingRepository) ListByOrganizer(ctx context.Context, organizerID string, page, pageSize int) ([]*entities.Meeting, int64, error) {
	offset := (page - 1) * pageSize

	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM meetings WHERE organizer_id = $1`
	if err := r.db.QueryRowContext(ctx, countQuery, organizerID).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Get meetings
	query := `
		SELECT id, room_id, title, description, organizer_id, start_time, end_time, status, jitsi_room_url, recording_url, max_participants, created_at, updated_at
		FROM meetings
		WHERE organizer_id = $1
		ORDER BY start_time DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, organizerID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	meetings := make([]*entities.Meeting, 0)
	for rows.Next() {
		meeting := &entities.Meeting{}
		var endTime sql.NullTime
		var jitsiURL, recordingURL sql.NullString

		if err := rows.Scan(
			&meeting.ID,
			&meeting.RoomID,
			&meeting.Title,
			&meeting.Description,
			&meeting.OrganizerID,
			&meeting.StartTime,
			&endTime,
			&meeting.Status,
			&jitsiURL,
			&recordingURL,
			&meeting.MaxParticipants,
			&meeting.CreatedAt,
			&meeting.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}

		if endTime.Valid {
			meeting.EndTime = &endTime.Time
		}
		if jitsiURL.Valid {
			meeting.JitsiRoomURL = jitsiURL.String
		}
		if recordingURL.Valid {
			meeting.RecordingURL = recordingURL.String
		}

		meetings = append(meetings, meeting)
	}

	return meetings, total, nil
}

// Update updates a meeting
func (r *MeetingRepository) Update(ctx context.Context, meeting *entities.Meeting) error {
	query := `
		UPDATE meetings
		SET title = $2, description = $3, start_time = $4, end_time = $5, status = $6, jitsi_room_url = $7, recording_url = $8, updated_at = $9
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		meeting.ID,
		meeting.Title,
		meeting.Description,
		meeting.StartTime,
		meeting.EndTime,
		meeting.Status,
		meeting.JitsiRoomURL,
		meeting.RecordingURL,
		meeting.UpdatedAt,
	)

	return err
}

// Delete deletes a meeting
func (r *MeetingRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM meetings WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetUpcoming retrieves upcoming meetings
func (r *MeetingRepository) GetUpcoming(ctx context.Context, userID string, limit int) ([]*entities.Meeting, error) {
	query := `
		SELECT id, room_id, title, description, organizer_id, start_time, end_time, status, jitsi_room_url, recording_url, max_participants, created_at, updated_at
		FROM meetings
		WHERE status IN ('scheduled', 'ongoing')
		ORDER BY start_time ASC
		LIMIT $1
	`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meetings := make([]*entities.Meeting, 0)
	for rows.Next() {
		meeting := &entities.Meeting{}
		var endTime sql.NullTime
		var jitsiURL, recordingURL sql.NullString

		if err := rows.Scan(
			&meeting.ID,
			&meeting.RoomID,
			&meeting.Title,
			&meeting.Description,
			&meeting.OrganizerID,
			&meeting.StartTime,
			&endTime,
			&meeting.Status,
			&jitsiURL,
			&recordingURL,
			&meeting.MaxParticipants,
			&meeting.CreatedAt,
			&meeting.UpdatedAt,
		); err != nil {
			return nil, err
		}

		if endTime.Valid {
			meeting.EndTime = &endTime.Time
		}
		if jitsiURL.Valid {
			meeting.JitsiRoomURL = jitsiURL.String
		}
		if recordingURL.Valid {
			meeting.RecordingURL = recordingURL.String
		}

		meetings = append(meetings, meeting)
	}

	return meetings, nil
}
