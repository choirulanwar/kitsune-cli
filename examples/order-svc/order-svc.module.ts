import { Module } from "@nestjs/common";
import { ClientsModule, Transport } from "@nestjs/microservices";
import { join } from "path";
import { ConsulService } from "src/common/consul.service";
import { protobufPackage as OrderSvcPackageName } from "__generated__/proto/order";
import { OrderService } from "./internal/core/service/service";
import { GqlHandler } from "./internal/handlers/gql";

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: "ORDER_SVC",
        inject: [ConsulService],
        useFactory: async (consul: ConsulService<any>) => {
          const svcName = "svc-order";
          const svcAddress = await consul.getService(svcName);

          return {
            transport: Transport.GRPC,
            options: {
              url: svcAddress?.url,
              package: OrderSvcPackageName,
              protoPath: join(process.cwd(), "./order.proto"),
              loader: {
                defaults: true,
                keepCase: true,
                longs: Number,
                enums: String,
                bytes: String,
                arrays: true,
                objects: true,
                oneofs: true,
                json: true,
                includeDirs: [join(process.cwd(), "../proto")],
              },
            },
          };
        },
      },
    ]),
  ],
  providers: [OrderService, GqlHandler],
  exports: [OrderService],
})
export class OrderModule {}
