import { User } from '@/model/common'

export interface Feedback {
  id: number;
  title: string;
  content: string;
  creator: number;
  creator_user?: User;
  created_at?: string;
  updated_at?: string;
}
