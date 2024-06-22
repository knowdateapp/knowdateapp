import { Workspace } from 'shared/model';

const defaultKey = 'notesAll';
export const queriesKeys = {
  all: defaultKey,
  notes: (workspace: Workspace) => [`${workspace}/notes`, defaultKey] as const,
};
