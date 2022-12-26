import { Inject, OnModuleInit } from "@nestjs/common";
import { ClientGrpc } from "@nestjs/microservices";
import { connectionFromArraySlice } from "graphql-relay";
import { lastValueFrom, map } from "rxjs";
import { getPagingParameters } from "src/common/relay/connection.args";
import { Order } from "../domain/order";
import { OrderBy, OrderType } from "__generated__/proto/common";
import { OrderService as OrderPbService } from "__generated__/proto/order";
import { StoreOrderInput } from "../../dto/store-order.input";

export class OrderService implements OnModuleInit {
  private orderService: OrderPbService;

  constructor(@Inject("ORDER_SVC") private client: ClientGrpc) {}

  onModuleInit() {
    this.orderService = this.client.getService<OrderPbService>("OrderService");
  }

  async Find(id: string): Promise<Order> {
    return await lastValueFrom(
      this.orderService.FindOrder({ id }).pipe(map(({ order }) => order))
    );
  }

  async Store(order: StoreOrderInput): Promise<Order> {
    const newStore = await lastValueFrom(
      this.orderService.StoreOrder({
        order: order,
      })
    );

    return this.Find(newStore.id);
  }

  async Update(id: string, order): Promise<Order> {
    return lastValueFrom(this.orderService.UpdateOrder({ id, order }));
  }

  async FindAll(args): Promise<any> {
    const { limit, offset, orderBy, orderType } = getPagingParameters(args);

    const data = await lastValueFrom(
      this.orderService.FindAllOrder({
        limit,
        offset,
        orderBy: OrderBy[orderBy.toUpperCase()],
        orderType: OrderType[orderType.toUpperCase()],
      })
    );

    const page = connectionFromArraySlice(data.orders, args, {
      arrayLength: data.count,
      sliceStart: offset || 0,
    });
    return { page, pageData: { count: data.count, limit, offset } };
  }

  async Delete(id: string): Promise<Order> {
    const order = await this.Find(id);

    await lastValueFrom(this.orderService.DeleteOrder({ id }));

    return order;
  }
}
