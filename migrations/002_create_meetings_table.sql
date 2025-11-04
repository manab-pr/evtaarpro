-- Create meetings table
CREATE TABLE IF NOT EXISTS meetings (
    id VARCHAR(36) PRIMARY KEY,
    room_id VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    organizer_id VARCHAR(36) NOT NULL REFERENCES users(id),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    status VARCHAR(20) NOT NULL CHECK (status IN ('scheduled', 'ongoing', 'completed', 'cancelled')),
    jitsi_room_url TEXT,
    recording_url TEXT,
    max_participants INT DEFAULT 50,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create meeting_participants table
CREATE TABLE IF NOT EXISTS meeting_participants (
    meeting_id VARCHAR(36) REFERENCES meetings(id) ON DELETE CASCADE,
    user_id VARCHAR(36) REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP,
    left_at TIMESTAMP,
    role VARCHAR(20) NOT NULL CHECK (role IN ('host', 'participant', 'guest')),
    PRIMARY KEY (meeting_id, user_id)
);

-- Create indexes
CREATE INDEX idx_meetings_organizer_id ON meetings(organizer_id);
CREATE INDEX idx_meetings_start_time ON meetings(start_time);
CREATE INDEX idx_meetings_status ON meetings(status);
CREATE INDEX idx_meeting_participants_user_id ON meeting_participants(user_id);

-- Create trigger for meetings table
CREATE TRIGGER update_meetings_updated_at BEFORE UPDATE ON meetings
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
