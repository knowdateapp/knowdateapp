import { apiClient } from 'shared/api';
import { Workspace } from 'shared/model';
import { CreateNoteParams, Note } from '../types';

const getUrl = (workspace: Workspace) => `${workspace}/notes`;

export const controller = {
  createNote: ({ workspace, title = null }: CreateNoteParams) =>
    apiClient.post<Note>(getUrl(workspace), { title }),
};
