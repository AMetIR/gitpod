/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the Gitpod Enterprise Source Code License,
 * See License.enterprise.txt in the project root folder.
 */

import { injectable, inject } from "inversify";
import { Repository } from "typeorm";
import { v4 as uuidv4 } from "uuid";

import { TypeORM } from "./typeorm";
import { WebhookEventDB } from "../webhook-event-db";
import { DBWebhookEvent } from "./entity/db-webhook-event";
import { WebhookEvent } from "@gitpod/gitpod-protocol";

@injectable()
export class WebhookEventDBImpl implements WebhookEventDB {
    @inject(TypeORM) protected readonly typeORM: TypeORM;

    protected async getEntityManager() {
        return (await this.typeORM.getConnection()).manager;
    }

    protected async getRepo(): Promise<Repository<DBWebhookEvent>> {
        return (await this.getEntityManager()).getRepository(DBWebhookEvent);
    }

    async createEvent(parts: Omit<WebhookEvent, "id">): Promise<WebhookEvent> {
        const repo = await this.getRepo();
        const newEvent: WebhookEvent = {
            ...parts,
            id: uuidv4(),
            creationTime: new Date().toISOString(),
        };
        return await repo.save(newEvent);
    }

    async updateEvent(id: string, update: Partial<WebhookEvent>): Promise<void> {
        const repo = await this.getRepo();
        const safeUpdate: Partial<WebhookEvent> = { ...update, id };
        delete safeUpdate.type;
        delete safeUpdate.rawEvent;
        await repo.save(safeUpdate);
    }

    async findByCloneUrl(cloneUrl: string, limit?: number): Promise<WebhookEvent[]> {
        const repo = await this.getRepo();
        const query = repo.createQueryBuilder("event");
        query.where("event.cloneUrl = :cloneUrl", { cloneUrl });
        query.orderBy("creationTime", "DESC");
        query.limit(limit);
        return query.getMany();
    }
}
