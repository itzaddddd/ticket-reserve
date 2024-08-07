ALTER TABLE events
ADD CONSTRAINT remain_quota_nonnegative CHECK (remain_quota >= 0)