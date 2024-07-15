import { Workspace } from 'shared/model';

// TODO: Что такое defaultKey?
const defaultKey = 'notes' as const;
export const queriesKeys = {
  all: defaultKey,
  notes: (workspace: Workspace) => [`${workspace}/notes`, defaultKey, workspace] as const,
  cards: (workspace: Workspace) => [`${workspace}/cards`, defaultKey, workspace] as const,
};
