import PageTitle from "../components/molecules/PageTitle";
import { getUser } from "../services/auth";

export default function ProfilePage() {
  const user = getUser();

  return (
    <div className="space-y-5">
      <PageTitle
        title="Profil"
        description="Informasi akun yang sedang login di browser ini."
      />

      <div className="grid gap-4 md:grid-cols-2">
        <section className="rounded-2xl border border-slate-200 bg-slate-50 p-4 shadow-sm">
          <p className="text-xs font-semibold uppercase tracking-wide text-slate-500">
            Username
          </p>
          <p className="mt-2 text-lg font-semibold text-slate-900">
            {user?.username ?? "-"}
          </p>
        </section>

        <section className="rounded-2xl border border-slate-200 bg-slate-50 p-4 shadow-sm">
          <p className="text-xs font-semibold uppercase tracking-wide text-slate-500">
            Role
          </p>
          <p className="mt-2 text-lg font-semibold text-slate-900">
            {user?.role ?? "-"}
          </p>
        </section>
      </div>
    </div>
  );
}
