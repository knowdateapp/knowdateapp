import { Workspace } from '../../../shared/model';
import { apiClient } from '../../../shared/api';
import { ICard } from '../types';
import { ICreateCardParams } from './types';

const getUrl = (workspace: Workspace) => `${workspace}/cards`;

export const controller = {
  createCard: ({ workspace, question, answer }: ICreateCardParams) =>
    apiClient.post<ICard>(getUrl(workspace), { question, answer }),
};
