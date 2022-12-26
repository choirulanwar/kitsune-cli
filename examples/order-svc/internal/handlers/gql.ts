import { UseGuards } from "@nestjs/common";
import { Args, Mutation, Query, Resolver } from "@nestjs/graphql";
import { Order } from "../core/domain/domain";
import { GqlUser } from "src/auth/internal/decorator";
import { GqlAuthGuard } from "src/auth/internal/graphql-auth.guard";
import ConnectionArgs from "src/common/relay/connection.args";
import { OrderService } from "../core/service/service";
import { FindAllOrderPayload } from "../dto/find-all.payload";
import { StoreOrderInput } from "../dto/store.input";
import { UpdateOrderInput } from "../dto/update.input";

@Resolver(() => Order)
export class GqlHandler {
  constructor(protected orderService: OrderService) {}

  @UseGuards(GqlAuthGuard)
  @Query(() => Order)
  async order(@Args("id") id: string): Promise<Order> {
    return this.orderService.Find(id);
  }

  @UseGuards(GqlAuthGuard)
  @Mutation(() => Order)
  async storeOrder(
    @Args("input") data: StoreOrderInput,
    @GqlUser() user
  ): Promise<Order> {
    return this.orderService.Store(data, user);
  }

  @UseGuards(GqlAuthGuard)
  @Mutation(() => Order)
  async updateOrder(
    @Args("id") id: string,
    @Args("input") data: UpdateOrderInput,
    @GqlUser() user
  ): Promise<Order> {
    return this.orderService.Update(id, data, user);
  }

  @UseGuards(GqlAuthGuard)
  @Query(() => FindAllOrderPayload)
  async orders(@Args("args") args: ConnectionArgs): Promise<any> {
    return this.orderService.FindAll(args);
  }

  @UseGuards(GqlAuthGuard)
  @Mutation(() => Order)
  async deleteOrder(@Args("id") id: string, @GqlUser() user): Promise<Order> {
    return this.orderService.Delete(id, user);
  }
}
