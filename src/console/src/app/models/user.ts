export type Role = 'Admin' | 'User';

export interface UserDetails extends User {
  role: Role;
}

export interface User {
  Id: string;
  Name?: string;
  username?: string;
  Email?: string;
  Phone?: string;
}
