import { useState } from "react";
import Swal from "sweetalert2";
import Button from "../components/atoms/Button";
import TextInput from "../components/atoms/TextInput";
import FormField from "../components/molecules/FormField";
import PageTitle from "../components/molecules/PageTitle";
import { changePassword } from "../services/auth";

const INITIAL_FORM = {
  oldPassword: "",
  newPassword: "",
  confirmPassword: "",
};

export default function PasswordPage() {
  const [form, setForm] = useState(INITIAL_FORM);
  const [errors, setErrors] = useState({});
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const validateForm = () => {
    const nextErrors = {};

    if (!form.oldPassword.trim()) {
      nextErrors.oldPassword = "Password lama wajib diisi";
    }

    if (!form.newPassword.trim()) {
      nextErrors.newPassword = "Password baru wajib diisi";
    }

    if (!form.confirmPassword.trim()) {
      nextErrors.confirmPassword = "Konfirmasi password wajib diisi";
    }

    if (
      form.newPassword &&
      form.confirmPassword &&
      form.newPassword !== form.confirmPassword
    ) {
      nextErrors.confirmPassword = "Konfirmasi password tidak sama";
    }

    setErrors(nextErrors);
    return Object.keys(nextErrors).length === 0;
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (!validateForm()) return;

    try {
      setLoading(true);
      setError("");
      await changePassword({
        oldPassword: form.oldPassword,
        newPassword: form.newPassword,
        confirmPassword: form.confirmPassword,
      });

      setForm(INITIAL_FORM);
      await Swal.fire({
        title: "Berhasil",
        text: "Password berhasil diubah.",
        icon: "success",
        confirmButtonText: "OK",
      });
    } catch (err) {
      setError(err.message);
      await Swal.fire({
        title: "Gagal",
        text: err.message || "Gagal mengubah password.",
        icon: "error",
        confirmButtonText: "Tutup",
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-5">
      <PageTitle
        title="Ubah Password"
        description="Gunakan password lama untuk membuat password baru."
      />

      {error && <p className="text-sm text-red-500">Error: {error}</p>}

      <form onSubmit={handleSubmit} className="max-w-xl space-y-4">
        <FormField
          label="Password Lama"
          htmlFor="oldPassword"
          error={errors.oldPassword}
        >
          <TextInput
            id="oldPassword"
            name="oldPassword"
            type="password"
            value={form.oldPassword}
            onChange={handleChange}
            autoComplete="current-password"
          />
        </FormField>

        <FormField
          label="Password Baru"
          htmlFor="newPassword"
          error={errors.newPassword}
        >
          <TextInput
            id="newPassword"
            name="newPassword"
            type="password"
            value={form.newPassword}
            onChange={handleChange}
            autoComplete="new-password"
          />
        </FormField>

        <FormField
          label="Konfirmasi Password Baru"
          htmlFor="confirmPassword"
          error={errors.confirmPassword}
        >
          <TextInput
            id="confirmPassword"
            name="confirmPassword"
            type="password"
            value={form.confirmPassword}
            onChange={handleChange}
            autoComplete="new-password"
          />
        </FormField>

        <Button type="submit" disabled={loading}>
          {loading ? "Menyimpan..." : "Ubah Password"}
        </Button>
      </form>
    </div>
  );
}
